package payments

import "errors"

var (
	errNilRequest = errors.New("nil request")
	errEmptyID    = errors.New("empty ID")
)

func validateGetPaymentRequest(request *GetPaymentRequest) error {
	if request == nil {
		return errNilRequest
	}
	
	if request.ID == "" {
		return errEmptyID
	}
	return nil
}
