package main

import (
	"fmt"
	"net/http"
)

type addHeaderHandler struct {
	handler   http.Handler
	headerMsg string
}

func (h *addHeaderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Hoge-Header", "golang")
	fmt.Println("HOGEHOGE")
	h.handler.ServeHTTP(w, r)

}

func newAddHeaderHandler(handler http.Handler, msg string) *addHeaderHandler {
	return &addHeaderHandler{
		handler:   handler,
		headerMsg: msg,
	}
}

type msgHandler struct {
	msg string
}

func (h *msgHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, h.msg)
}

func main() {
	http.Handle("/", newAddHeaderHandler(&msgHandler{"Body message"}, "header message!!!"))

	http.ListenAndServe(":8080", nil)
}
