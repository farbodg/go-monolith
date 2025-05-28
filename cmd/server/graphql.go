package server

import (
	"github.com/99designs/gqlgen/graphql"
	"go-monolith/api/graphql/graph"
	"go-monolith/api/graphql/graph/resolvers"
)

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return graph.NewExecutableSchema(graph.Config{
		Resolvers: s.GQLResolvers,
	})
}

func (s *Server) Mutation() graph.MutationResolver {
	return &resolvers.Mutation{}
}

func (s *Server) Query() graph.QueryResolver {
	return &resolvers.Query{
		AccountsResolver: &resolvers.AccountsResolver{
			AccountsService: s.Services.AccountsService,
		},
		PaymentsResolver: &resolvers.PaymentsResolver{
			PaymentsService: s.Services.PaymentsService,
		},
	}
}
