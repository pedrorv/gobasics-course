package main

import (
	"fmt"
	"net/http"
)

type custom int

func (c custom) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Something")
}

func main() {
	var c custom
	http.ListenAndServe(":8080", c)
}