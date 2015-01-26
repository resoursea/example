package main

import (
	"bytes"
	"encoding/json"
	"github.com/dchest/uniuri"
	"net/http"
	"reflect"
	"testing"
)

//
// Test BookPOST
//
func TestBookPOST(t *testing.T) {

	book := &BookPOST{
		Title:       "New Book - " + uniuri.NewLen(5),
		Description: "New description",
	}

	b, err := json.Marshal(book)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("POST", "/api/categories/animals/books", bytes.NewBuffer(b))
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/json")

	bookResp, err := book.Init(req)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(book, bookResp) {
		t.Error("Book returned by BookPOST is wrong")
	}

}
