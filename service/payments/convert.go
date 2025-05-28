package payments

import (
	"go-monolith/service/payments/data"
)

func paymentDALToService(paymentDAL *data.Payment) *Payment {
	if paymentDAL == nil {
		return nil
	}

	return &Payment{
		ID:          paymentDAL.ID,
		Amount:      paymentDAL.Amount,
		ReferenceID: paymentDAL.ReferenceID,
		AccountID:   paymentDAL.AccountID,
		CreatedAt:   paymentDAL.CreatedAt,
		UpdatedAt:   paymentDAL.UpdatedAt,
	}
}
