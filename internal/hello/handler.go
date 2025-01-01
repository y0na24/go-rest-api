package hello

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func NewHelloHandler(mux *http.ServeMux) {
	handler := &HelloHandler{}
	mux.HandleFunc("/hello", handler.Hello())
}

func (handler *HelloHandler) Hello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello")
	}
}
