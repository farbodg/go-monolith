package payments

import "time"

type GetPaymentRequest struct {
	ID string
}

type GetPaymentResponse struct {
	Payment *Payment
}

type Payment struct {
	ID          string
	Amount      string
	ReferenceID string
	AccountID   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
