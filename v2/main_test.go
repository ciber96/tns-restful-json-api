package main

import(
	"testing"
	"net/http"
	"io/ioutil"
	"log"
)


func TestMain(t *testing.T) {
	resp, err := http.Get("/")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(body))
}