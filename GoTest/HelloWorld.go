package main

import (
	"fmt"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, %s !", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}