package util

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code  int
	Error error
}

func ErrorHandle(code int, err error) *Error {
	fmt.Printf("%+v", err)
	return &Error{
		Code:  code,
		Error: err,
	}
}

func DefaultErrorHandle(err error) *Error {
	fmt.Printf("%+v", err)
	return &Error{
		Code:  http.StatusInternalServerError,
		Error: err,
	}
}

// func ErrorMessage(err error, msg string) error {
// 	fmt.Printf("%+v", err)
// 	errWrap := errors.Wrap(err, msg)
// 	return errWrap
// }
