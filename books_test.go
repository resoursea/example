package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type BooksResp struct {
	Book []Book
}

//
// Test [GET] /api/categories/uncategorized/books
//
func TestGETBooks(t *testing.T) {

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/categories/uncategorized/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	route.ServeHTTP(w, req)

	// Check if any error was returned
	errData := &ErrorResp{}
	err = json.Unmarshal(w.Body.Bytes(), errData)
	if err != nil {
		t.Fatal(err)
	}
	if len(errData.Error) > 0 {
		t.Fatal(errData.Error)
	}

	// Try to get the book from the response
	booksResp := &BooksResp{}
	err = json.Unmarshal(w.Body.Bytes(), booksResp)
	if err != nil {
		t.Fatal(err)
	}

	if len(booksResp.Book) == 1 {
		t.Fatal("Book list not returned into the answare")
	}
}
