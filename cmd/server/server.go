package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/oklog/run"
	"go-monolith/api/graphql/graph"
	"go-monolith/service/accounts"
	"go-monolith/service/payments"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	GQLResolvers  graph.ResolverRoot
	RouterHandler http.Handler
	Port          uint16
	Services      Services
}

func NewServer(conf Config) *Server {
	server := &Server{
		Port: conf.Port,
	}

	server.GQLResolvers = server
	server.Services = Services{
		AccountsService: accounts.New(accounts.Config{DB: conf.DB}),
		PaymentsService: payments.New(payments.Config{DB: conf.DB}),
	}

	router := SetupRouter(server)
	server.RouterHandler = router

	return server
}

func (s *Server) Run(ctx context.Context, g *run.Group) error {
	s.createHTTPServer(ctx, g)
	s.listenToInterrupt(g)
	return nil
}

func SetupRouter(s *Server) http.Handler {
	r := chi.NewRouter()

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Timeout(30*time.Second),
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	r.Post("/api/graphql", graphQLHandler(s))
	r.Get("/api/graphql/playground", graphQLPlaygroundHandler())

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	return r
}

func graphQLHandler(s *Server) http.HandlerFunc {
	h := handler.New(s.ToExecutableSchema())

	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})
	h.Use(extension.Introspection{})

	return h.ServeHTTP
}

func graphQLPlaygroundHandler() http.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/api/graphql")
	return h.ServeHTTP
}

func (s *Server) createHTTPServer(ctx context.Context, g *run.Group) {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.Port),
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		Handler:      s.RouterHandler,
	}

	g.Add(
		func() error {
			fmt.Printf("HTTP server running on port %d\n", s.Port)
			return server.ListenAndServe()
		},
		func(err error) {
			shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			_ = server.Shutdown(shutdownCtx)
		},
	)
}

func (s *Server) listenToInterrupt(g *run.Group) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	doneCancel := make(chan struct{})

	g.Add(
		func() error {
			select {
			case sig := <-sigChan:
				fmt.Println("Service Termination: ", sig.String())
				if sig == os.Interrupt {
					return errors.New("OS INTERRUPT SIGNAL RECEIVED")
				}
			case <-doneCancel:
			}
			return nil
		},
		func(e error) {
			close(doneCancel)
		},
	)
}
