package main

import (
	"context"
	"errors"
	"fmt"
	"go-monolith/cmd/server"
	configpkg "go-monolith/config"
	dbpkg "go-monolith/db"
	"os"
	"os/signal"

	"github.com/kelseyhightower/envconfig"
	"github.com/oklog/run"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	var config configpkg.Config
	var g run.Group

	if err := envconfig.Process("", &config); err != nil {
		fmt.Printf("falied to load envvars into config: %v\n", err)
		os.Exit(1)
	}

	if err := config.ComputeDependencies(); err != nil {
		fmt.Printf("falied to load dependencies into config: %v\n", err)
		os.Exit(1)
	}

	db, err := dbpkg.ConnectToDB(ctx, config.DBConfig)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v\n", err)
		os.Exit(1)
	}

	if err := dbpkg.Migrate(db.GetDB().DB); err != nil {
		fmt.Printf("failed to migrate database: %v\n", err)
		os.Exit(1)
	}

	serverImpl := server.NewServer(server.Config{Port: config.ServerHTTPPort, DB: db})
	if err := serverImpl.Run(ctx, &g); err != nil {
		fmt.Printf("failed to start server: %v\n", err)
		os.Exit(1)
	}

	g.Add(
		func() error {
			<-ctx.Done()
			return ctx.Err()
		},
		func(err error) {
			cancel()
			db.Close()
		},
	)

	if err := g.Run(); err != nil && !errors.Is(err, context.Canceled) {
		fmt.Println("exited with error:", err)
	}
}
