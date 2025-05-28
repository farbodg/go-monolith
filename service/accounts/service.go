package accounts

import (
	"context"
	"fmt"
	"go-monolith/service/accounts/data"
)

type Service interface {
	GetAccount(ctx context.Context, request *GetAccountRequest) (*GetAccountResponse, error)
}

type impl struct {
	dal data.DALProvider
}

func New(conf Config) Service {
	return &impl{
		dal: data.New(conf.DB),
	}
}

func (i *impl) GetAccount(ctx context.Context, request *GetAccountRequest) (*GetAccountResponse, error) {
	if err := validateGetAccountRequest(request); err != nil {
		return nil, err
	}

	accountDAL := &data.Account{ID: request.ID}

	if err := i.dal.GetAccountByID(ctx, accountDAL); err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	return &GetAccountResponse{
		Account: accountDALToService(accountDAL),
	}, nil
}
