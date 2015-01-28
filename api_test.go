package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// The Api Resource response for [GET] /api
type ApiData struct {
	Api ApiResp
}

type ApiResp struct {
	Version int
	Welcome string
}

//
// Test [GET] /api
//
func TestApi(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api", nil)
	if err != nil {
		t.Fatal(err)
	}

	route.ServeHTTP(w, req)

	var resp ApiData
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Api.Version != 1 && resp.Api.Welcome != "This is the REST API for a book store" {
		t.Fatal("[GET] /api answare is something different!")
	}

}
