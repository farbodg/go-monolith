package payments

import (
	"go-monolith/service/payments/data"
	"reflect"
	"testing"
	"time"
)

func Test_PaymentDALToService(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name       string
		input      *data.Payment
		wantOutput *Payment
	}{
		{
			name: "Normal mapping",
			input: &data.Payment{
				ID:          "pay-1",
				Amount:      "42.50",
				ReferenceID: "ref-abc",
				AccountID:   "account-id",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantOutput: &Payment{
				ID:          "pay-1",
				Amount:      "42.50",
				ReferenceID: "ref-abc",
				AccountID:   "account-id",
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		},
		{
			name:       "Nil input",
			input:      nil,
			wantOutput: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := paymentDALToService(tt.input)
			if !reflect.DeepEqual(got, tt.wantOutput) {
				t.Errorf("paymentDALToService(%+v) = %+v, want %+v", tt.input, got, tt.wantOutput)
			}
		})
	}
}
