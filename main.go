package main

import (
    "fmt"
	"net/http"
    "text/template"
)

type Page struct {
    Title string
    Count int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    page := Page{"Hello World.", 1}
    file := "layout.html"
    tmpl, err := template.ParseFiles(file)
    if err != nil {
        panic(err)
    }

    err = tmpl.Execute(w, page)
    if err != nil {
        panic(err)
    }
}

func main() {
     http.HandleFunc("/", viewHandler) // hello
	fmt.Println("Server Start")
	http.ListenAndServe(":8080", nil)
}
