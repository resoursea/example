package main

import (
	"encoding/json"
	"net/http"
)

// This resource gets the data sent by the client when saving a new book
type BookPOST struct {
	Title       string
	Description string
}

// This method is called when this resource is requested
// It retrieves the BookPOST struct contained in the request body
// and return it to the requester method
func (b *BookPOST) Init(req *http.Request) (*BookPOST, error) {
	err := json.NewDecoder(req.Body).Decode(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
