package main

import(
	"testing"
	"net/http"
	"io/ioutil"
)


func TestMain(t *testing.T) {
	resp, err := http.Get("/")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.body)
	if err != nil {
		t.Fatal(err)
	}
}