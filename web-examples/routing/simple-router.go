package main

import (
	"io"
	"net/http"
)

type server int

func (s server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		io.WriteString(res, "Hello world!")
	case "/secret":
		io.WriteString(res, "Secret message.")
	}
}

func main() {
	var s server
	http.ListenAndServe(":8080", s)
}
