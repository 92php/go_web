package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

//http://127.0.0.1/hello
func main() {
	server := &http.Server{
		Addr: "0.0.0.0:80",
	}
	http.HandleFunc("/", hello)
	server.ListenAndServe()
}
