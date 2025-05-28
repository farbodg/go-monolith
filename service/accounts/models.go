package accounts

import "time"

type GetAccountRequest struct {
	ID string
}

type GetAccountResponse struct {
	Account *Account
}

type Account struct {
	ID        string
	AuthID    string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
