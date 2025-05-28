package accounts

import (
	"errors"
	"testing"
)

func Test_ValidateGetAccountRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     *GetAccountRequest
		wantErr error
	}{
		{
			name:    "Valid ID",
			req:     &GetAccountRequest{ID: "123"},
			wantErr: nil,
		},
		{
			name:    "Empty ID",
			req:     &GetAccountRequest{ID: ""},
			wantErr: errEmptyID,
		},
		{
			name:    "Nil request",
			req:     nil,
			wantErr: errNilRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateGetAccountRequest(tt.req)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got error %v, want %v", err, tt.wantErr)
			}
		})
	}
}
