package errors

import (
	"errors"
	"fmt"
)

/*
Here we defined 2 error data-types one for bad_request and another for authentication
*/
var (
	ErrBadRequest     = errors.New("bad request")
	ErrAuthentication = errors.New("authentication failed")
)

func WebCall(b bool) error {
	if b {
		return ErrBadRequest
	}
	return ErrAuthentication
}

func WebCallSwitchCase(b bool) error {
	if err := WebCall(b); err != nil {
		switch {
		case errors.Is(err, ErrBadRequest):
			fmt.Println("Bad request happened.")
			return ErrBadRequest
		case errors.Is(err, ErrAuthentication):
			fmt.Println("Authentication happened.")
			return ErrAuthentication
		}
	}

	return nil
}
