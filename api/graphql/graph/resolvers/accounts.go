package resolvers

import (
	"context"
	"go-monolith/api/graphql/graph/models"
	"go-monolith/service/accounts"
)

type AccountsResolver struct {
	AccountsService accounts.Service
}

func (r *AccountsResolver) Account(ctx context.Context, input models.AccountInput) (*models.AccountResponse, error) {
	accountRes, err := r.AccountsService.GetAccount(ctx, &accounts.GetAccountRequest{ID: input.ID})
	if err != nil {
		return nil, err
	}

	return &models.AccountResponse{
		Account: accountServiceToGQL(accountRes.Account),
	}, nil
}
