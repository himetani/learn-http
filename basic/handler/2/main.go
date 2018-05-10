package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux
	fmt.Println(mux)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		http.HandleFunc("/append", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello appended endpoint")
		})
		fmt.Fprintf(w, "Hello world")
	})

	fmt.Println(mux)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
