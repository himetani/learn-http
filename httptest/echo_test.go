package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEcho(t *testing.T) {
	recorder := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "localhost", nil)
	if err != nil {
		t.Error("Unexpected error happend")
	}

	handler := echoHandler()
	handler(recorder, req)

	bytes, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Error("Unexpected error happend")
	}

	if string(bytes) != "echo" {
		t.Errorf("expected %s, buy got %s\n", "echo", string(bytes))
	}

}
