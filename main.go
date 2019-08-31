package main

import (
	"io"
	"net/http"
)

func brad(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Brad")
}

func march(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "March")
}

func main() {

	http.Handle("/Brad", http.HandlerFunc(brad))
	http.Handle("/March", http.HandlerFunc(march))

	http.ListenAndServe(":8080", nil)
}
