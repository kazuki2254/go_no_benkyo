//testserver

package main

import (
	"fmt"
	"log"
	"net/http"
)

//hogehogeほげほげ

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.WriteHeader(400)
		fmt.Fprintf(w, "URL=:%v\tcall GET", r.URL.Path)
	case "POST":
		fmt.Fprintf(w, "URL=:%v\tcall POST", r.URL.Path)
	case "PUT":
		fmt.Fprintf(w, "URL=:%v\tcall PUT", r.URL.Path)
	case "DELETE":
		fmt.Fprintf(w, "URL=:%v\tcall DELETE", r.URL.Path)
	default:
		fmt.Fprintf(w, "URL=:%v\tunknown call: %v", r.Method, r.URL.Path)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
