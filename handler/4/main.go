package main

import (
	"fmt"
	"net/http"
)

type hoge int

func (h *hoge) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hoge type: %d", *h)
}

func main() {
	mux := http.NewServeMux()
	h := hoge(9)
	mux.Handle("/hoge", &h)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
