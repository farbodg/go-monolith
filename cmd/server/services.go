package server

import (
	"go-monolith/service/accounts"
	"go-monolith/service/payments"
)

type Services struct {
	AccountsService accounts.Service
	PaymentsService payments.Service
}
