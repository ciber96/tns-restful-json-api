package main

import (
	"net/http"
	"testing"
	"net/http/httptest"
)

func TestGetEntries(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users", nil)
    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    if body := response.Body.String(); body != "Hello, / | Production Branch" {
        t.Errorf("Expected an empty array. Got %s", body)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    Router := mux.NewRouter()
    Router.ServeHTTP(rr, req)

    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}