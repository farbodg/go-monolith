package resolvers

import (
	"go-monolith/api/graphql/graph/models"
	"go-monolith/service/accounts"
	"go-monolith/service/payments"
)

func accountServiceToGQL(account *accounts.Account) *models.Account {
	return &models.Account{
		ID:        account.ID,
		AuthID:    account.AuthID,
		Email:     account.Email,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}

func paymentServiceToGQL(payment *payments.Payment) *models.Payment {
	return &models.Payment{
		ID:          payment.ID,
		Amount:      payment.Amount,
		ReferenceID: payment.ReferenceID,
		AccountID:   payment.AccountID,
		CreatedAt:   payment.CreatedAt,
		UpdatedAt:   payment.UpdatedAt,
	}
}
