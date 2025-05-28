package accounts

import "errors"

var (
	errNilRequest = errors.New("nil request")
	errEmptyID    = errors.New("empty ID")
)

func validateGetAccountRequest(req *GetAccountRequest) error {
	if req == nil {
		return errNilRequest
	}
	
	if req.ID == "" {
		return errEmptyID
	}
	return nil
}
