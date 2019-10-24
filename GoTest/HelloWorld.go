package main

import (
	"fmt"
	"net/http"
	"src/github.com/gorilla/mux"
)

func HelloServer(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, %s !", r.URL.Path[1:])
}


func main() {
	//http.HandleFunc("/", HelloServer)
	//
	//fs := http.FileServer(http.Dir("D:/IdeaProjects/StreamTest/GoTest/static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request)  {
		vars := mux.Vars(r)
		a := vars["title"] // the book title slug
		b := vars["page"] // the page
		fmt.Printf(a + ":" + b)
	})

	r.HandleFunc("/t1/*", func(w http.ResponseWriter, r *http.Request)  {
		fmt.Printf("t1")
	})
	r.HandleFunc("/t2/*", func(w http.ResponseWriter, r *http.Request)  {
		fmt.Printf("t2")
	})
	
	
	http.ListenAndServe(":8080", r)
}