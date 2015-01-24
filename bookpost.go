package main

import (
	"encoding/json"
	"net/http"
)

// This resource gets the data sent by user
// when posting a new book
type BookPOST struct {
	Title       string
	Description string
}

func (b *BookPOST) Init(req *http.Request) (*BookPOST, error) {
	err := json.NewDecoder(req.Body).Decode(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
