package accounts

import (
	"go-monolith/service/accounts/data"
)

func accountDALToService(accountDAL *data.Account) *Account {
	if accountDAL == nil {
		return nil
	}
	
	return &Account{
		ID:        accountDAL.ID,
		AuthID:    accountDAL.AuthID,
		Email:     accountDAL.Email,
		CreatedAt: accountDAL.CreatedAt,
		UpdatedAt: accountDAL.UpdatedAt,
	}
}
