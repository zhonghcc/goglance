package main

import (
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello "+req.URL.String())
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}
