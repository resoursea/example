package main

import (
	"fmt"
)

type ErrorResp struct {
	Error string
}

type ErrorsResp struct {
	Error []string
}

func (e *ErrorResp) String() string {
	return e.Error
}

func (e *ErrorsResp) String() string {
	return fmt.Sprintf("%q", e.Error)
}
