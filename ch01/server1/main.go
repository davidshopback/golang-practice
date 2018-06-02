package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "URL.Path = %q\n", req.URL.Path)
}
