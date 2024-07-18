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
	fmt.Printf("error! --> %+v", err)
	fmt.Println()
	return &Error{
		Code:  code,
		Error: err,
	}
}

func DefaultErrorHandle(err error) *Error {
	fmt.Printf("error! --> %+v", err)
	fmt.Println()
	return &Error{
		Code:  http.StatusInternalServerError,
		Error: err,
	}
}
