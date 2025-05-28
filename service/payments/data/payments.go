package data

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type Payment struct {
	bun.BaseModel `bun:"table:payments"`

	ID          string    `bun:"id,pk"`
	Amount      string    `bun:"amount,notnull"`
	ReferenceID string    `bun:"reference_id,notnull"`
	AccountID   string    `bun:"account_id,notnull"`
	CreatedAt   time.Time `bun:"created_at,notnull"`
	UpdatedAt   time.Time `bun:"updated_at,notnull"`
}

func (dp *dataProvider) GetPaymentByID(ctx context.Context, payment *Payment) error {
	return dp.GetDB().NewSelect().Model(payment).WherePK().Scan(ctx)
}
