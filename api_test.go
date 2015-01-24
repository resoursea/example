package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

//
// Test [GET] /api
//
func TestApi(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api", nil)
	if err != nil {
		t.Error(err)
	}

	route.ServeHTTP(w, req)

	var resp string
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Error(err)
	}

	if resp != "This is the REST API for a book store" {
		t.Error("[GET] /api answare is something different!")
	}
}
