package resolvers

import (
	"context"
	"go-monolith/api/graphql/graph/models"
	"go-monolith/service/payments"
)

type PaymentsResolver struct {
	PaymentsService payments.Service
}

func (r *PaymentsResolver) Payment(ctx context.Context, input models.PaymentInput) (*models.PaymentResponse, error) {
	paymentRes, err := r.PaymentsService.GetPayment(ctx, &payments.GetPaymentRequest{ID: input.ID})
	if err != nil {
		return nil, err
	}

	return &models.PaymentResponse{
		Payment: paymentServiceToGQL(paymentRes.Payment),
	}, nil
}
