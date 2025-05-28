package data

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts"`

	ID        string    `bun:"id,pk"`
	AuthID    string    `bun:"auth_id,notnull"`
	Email     string    `bun:"email,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull"`
	UpdatedAt time.Time `bun:"updated_at,notnull"`
}

func (dp *dataProvider) GetAccountByID(ctx context.Context, account *Account) error {
	return dp.GetDB().NewSelect().Model(account).WherePK().Scan(ctx)
}
