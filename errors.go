package main

// This struct are used to unmarshar the Error from the response
type ErrorResp struct {
	Error string
}

func (e *ErrorResp) String() string {
	return e.Error
}
