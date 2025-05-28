package payments

import (
	"context"
	"fmt"
	"go-monolith/service/payments/data"
)

type Service interface {
	GetPayment(ctx context.Context, request *GetPaymentRequest) (*GetPaymentResponse, error)
}

type impl struct {
	dal data.DALProvider
}

func New(conf Config) Service {
	return &impl{
		dal: data.New(conf.DB),
	}
}

func (i *impl) GetPayment(ctx context.Context, request *GetPaymentRequest) (*GetPaymentResponse, error) {
	if err := validateGetPaymentRequest(request); err != nil {
		return nil, err
	}

	paymentDAL := &data.Payment{ID: request.ID}

	if err := i.dal.GetPaymentByID(ctx, paymentDAL); err != nil {
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}

	return &GetPaymentResponse{
		Payment: paymentDALToService(paymentDAL),
	}, nil
}
