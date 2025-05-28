package resolvers

import (
	"go-monolith/api/graphql/graph/models"
	"go-monolith/service/accounts"
	"go-monolith/service/payments"
	"reflect"
	"testing"
	"time"
)

func Test_AccountServiceToGQL(t *testing.T) {
	now := time.Now()
	serviceAccount := &accounts.Account{
		ID:        "id-123",
		AuthID:    "auth-456",
		Email:     "user@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	want := &models.Account{
		ID:        "id-123",
		AuthID:    "auth-456",
		Email:     "user@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	got := accountServiceToGQL(serviceAccount)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("accountServiceToGQL() = %+v, want %+v", got, want)
	}
}

func Test_PaymentServiceToGQL(t *testing.T) {
	now := time.Now()
	servicePayment := &payments.Payment{
		ID:          "pay-789",
		Amount:      "100.50",
		ReferenceID: "ref-xyz",
		AccountID:   "account-id",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	want := &models.Payment{
		ID:          "pay-789",
		Amount:      "100.50",
		ReferenceID: "ref-xyz",
		AccountID:   "account-id",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	got := paymentServiceToGQL(servicePayment)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("paymentServiceToGQL() = %+v, want %+v", got, want)
	}
}
