package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request, result int) {

}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/view.html")

	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	var num1 int
	var num2 int
	num1, _ = strconv.Atoi(r.FormValue("num1"))
	num2, _ = strconv.Atoi(r.FormValue("num2"))
	op := r.FormValue("op")
	if op == "add" {
		result := num1 + num2
		t, err := template.ParseFiles("template/view.html")
		if err != nil {
			log.Fatal(err)
		}
		if err := t.Execute(w, result); err != nil {
			log.Fatal(err)
		}
	} else if op == "sub" {
		result := num1 - num2
		t, err := template.ParseFiles("template/view.html")
		if err != nil {
			log.Fatal(err)
		}
		if err := t.Execute(w, result); err != nil {
			log.Fatal(err)
		}
	} else if op == "mult" {
		result := num1 * num2
		t, err := template.ParseFiles("template/view.html")
		if err != nil {
			log.Fatal(err)
		}
		if err := t.Execute(w, result); err != nil {
			log.Fatal(err)
		}
	} else {
		result := num1 / num2
		t, err := template.ParseFiles("template/view.html")
		if err != nil {
			log.Fatal(err)
		}
		if err := t.Execute(w, result); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/calc", calcHandler)
	http.ListenAndServe(":8080", nil)
}
