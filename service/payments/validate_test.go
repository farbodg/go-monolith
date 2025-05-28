package payments

import (
	"errors"
	"testing"
)

func Test_ValidateGetPaymentRequest(t *testing.T) {
	tests := []struct {
		name    string
		input   *GetPaymentRequest
		wantErr error
	}{
		{
			name:    "Valid request",
			input:   &GetPaymentRequest{ID: "payment-123"},
			wantErr: nil,
		},
		{
			name:    "Nil request",
			input:   nil,
			wantErr: errNilRequest,
		},
		{
			name:    "Empty ID",
			input:   &GetPaymentRequest{ID: ""},
			wantErr: errEmptyID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateGetPaymentRequest(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("validateGetPaymentRequest(%+v) = %v, want %v", tt.input, err, tt.wantErr)
			}
		})
	}
}
