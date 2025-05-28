package db

import (
	"context"

	"github.com/uptrace/bun"
)

type BaseDAL interface {
	GetDB() *bun.DB
	Close() error
	DoInTx(ctx context.Context, fn func(ctx context.Context) error) error
}
