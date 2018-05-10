package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func AdminFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "admin\n")
}

func NormalFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "normal\n")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/admin", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(AdminFunc)))
	mux.HandleFunc("/", NormalFunc)

	// Wrap our server with our gzip handler to gzip compress all responses.
	http.ListenAndServe(":8000", handlers.CompressHandler(mux))
}
