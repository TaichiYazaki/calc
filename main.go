package main

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func viewhandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/view.html")

	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", viewhandler)
	http.ListenAndServe(":8080", nil)
}
