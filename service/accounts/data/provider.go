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
	GetAccountByID(ctx context.Context, account *Account) error
}
