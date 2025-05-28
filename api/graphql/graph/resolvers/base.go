package resolvers

import (
	"context"
)

type Mutation struct{}

func (m *Mutation) NoOp(ctx context.Context) (*bool, error) {
	b := true
	return &b, nil
}

type Query struct {
	*AccountsResolver
	*PaymentsResolver
}

func (q Query) NoOp(ctx context.Context) (*bool, error) {
	b := true
	return &b, nil
}
