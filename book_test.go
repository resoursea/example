package main

import (
	"bytes"
	"encoding/json"
	"github.com/dchest/uniuri"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type BookResp struct {
	Book Book
}

//
// Test [GET] /api/categories/default/books/default-book
//
func TestGETBook(t *testing.T) {

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/categories/default/books/default-book", nil)
	if err != nil {
		log.Panic(err)
	}

	route.ServeHTTP(w, req)

	// Check if any error was returned
	errData := &ErrorResp{}
	err = json.Unmarshal(w.Body.Bytes(), errData)
	if err != nil {
		t.Error(err)
	}
	if len(errData.Error) > 0 {
		t.Error(errData.Error)
	}

	// Try to get the book from the response
	bookResp := &BookResp{}
	err = json.Unmarshal(w.Body.Bytes(), bookResp)
	if err != nil {
		t.Error(err)
	}

	if bookResp.Book.BookID != "default" {
		t.Error("Book default not returned into the answare")
	}
}

//
// Test [POST] /api/categories/default/books
//
func TestPOSTBook(t *testing.T) {

	bp := BookPOST{
		Title:       "New Book - " + uniuri.NewLen(5),
		Description: "New description",
	}

	b, err := json.Marshal(bp)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/categories/animals/books", bytes.NewBuffer(b))
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/json")

	route.ServeHTTP(w, req)

	// Check if any error was returned
	errData := &ErrorResp{}
	err = json.Unmarshal(w.Body.Bytes(), errData)
	if err != nil {
		t.Error(err)
	}
	if len(errData.Error) > 0 {
		t.Error(errData.Error)
	}

	// Try to get the book from the response
	bookResp := &BookResp{}
	err = json.Unmarshal(w.Body.Bytes(), bookResp)
	if err != nil {
		t.Error(err)
	}
	if len(bookResp.Book.BookID) == 0 {
		t.Error("Book not returned into the answare")
	}

}

//
// Test [POST] /api/categories/default/books/default-book/like
//
func TestPOSTBookLike(t *testing.T) {

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/categories/animals/books/default-book/like", nil)
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/json")

	route.ServeHTTP(w, req)

	// Check if any error was returned
	errData := &ErrorResp{}
	err = json.Unmarshal(w.Body.Bytes(), errData)
	if err != nil {
		t.Error(err)
	}
	if len(errData.Error) > 0 {
		t.Error(errData.Error)
	}

	// Try to get the book from the response
	bookResp := &BookResp{}
	err = json.Unmarshal(w.Body.Bytes(), bookResp)
	if err != nil {
		t.Error(err)
	}
	if len(bookResp.Book.BookID) == 0 {
		t.Error("Book not returned into the answare from the book like")
	}

}
