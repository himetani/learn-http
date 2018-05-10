package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hoge", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
