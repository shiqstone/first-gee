package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/hello", helloeHandle)
	log.Fatal(http.ListenAndServe(":9002", nil))
}

func indexHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
}

func helloeHandle(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}