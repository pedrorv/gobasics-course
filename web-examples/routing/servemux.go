package main

import (
	"io"
	"net/http"
)

type handler1 int
type handler2 int

func (h handler1) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Handler one.")
}

func (h handler2) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Handler two.")
}

func main() {
	var h1 handler1
	var h2 handler2

	mux := http.NewServeMux()
	mux.Handle("/one", h1)
	mux.Handle("/two", h2)

	http.ListenAndServe(":8080", mux)
}
