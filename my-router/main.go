package main

import (
	"net/http"

	"github.com/himetani/router"
)

func main() {
	r := router.New()

	http.ListenAndServe(":8080", r)
}
