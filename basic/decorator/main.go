package main

import (
	"fmt"
	"net/http"
)

func addHeaderHandler(headerMsg string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Hoge-Header", headerMsg)
		handler.ServeHTTP(w, r)
	})
}

func bodyWriteHandler(bodyMsg string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, bodyMsg)
	})
}

func main() {
	closure1 := bodyWriteHandler("Hello gophers")

	http.Handle("/", addHeaderHandler("golang", addHeaderHandler("My-Header", closure1)))
	http.ListenAndServe(":8080", nil)
}
