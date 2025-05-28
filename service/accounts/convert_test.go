package accounts

import (
	"go-monolith/service/accounts/data"
	"reflect"
	"testing"
	"time"
)

func Test_AccountDALToService(t *testing.T) {
	now := time.Now()
	accountDAL := &data.Account{
		ID:        "acc123",
		AuthID:    "auth456",
		Email:     "test@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	expected := &Account{
		ID:        "acc123",
		AuthID:    "auth456",
		Email:     "test@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	t.Run("standard mapping", func(t *testing.T) {
		got := accountDALToService(accountDAL)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %+v, want %+v", got, expected)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		got := accountDALToService(nil)
		if got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
	})
}
