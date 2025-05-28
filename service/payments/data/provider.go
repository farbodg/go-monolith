package data

import (
	"context"
	"go-monolith/db"
)

type DALProvider interface {
	db.BaseDAL
	EventsProvider
}

type EventsProvider interface {
	GetPaymentByID(ctx context.Context, account *Payment) error
}
