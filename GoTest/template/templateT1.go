package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos []Todo
}

func main()  {
	tmpl := template.Must(template.ParseFiles("../layout.html"))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := TodoPageData{
			PageTitle:"My TODO list",
			Todos:[]Todo{
				{Title:"Task 1", Done: false},
				{Title:"Task 2", Done: true},
				{Title:"Task 3", Done: true},
				{Title:"Task 4", Done: false},
			},
		}
		tmpl.Execute(writer, data)
	})
	
	http.ListenAndServe(":8080", nil)
}