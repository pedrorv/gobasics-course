package main

import (
	"io"
	"net/http"
)

func h1(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Handler one.")
}

func h2(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Handler two.")
}

func main() {
	http.HandleFunc("/one", h1)
	http.HandleFunc("/two", h2)

	http.ListenAndServe(":8080", nil)
}
