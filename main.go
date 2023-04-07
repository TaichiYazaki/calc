package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Calc struct {
	LeftNum  int
	RightNum int
	Operator string
	Result   int
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
	var cal Calc
	cal.Operator = r.FormValue("operator")
	cal.LeftNum, _ = strconv.Atoi(r.FormValue("leftNum"))
	cal.RightNum, _ = strconv.Atoi(r.FormValue("rightNum"))

	if cal.Operator == "add" {
		cal.Result = cal.LeftNum + cal.RightNum
		t, err := template.ParseFiles("template/view.html")
		if err != nil {
			log.Fatal(err)
		}
		if err := t.Execute(w, cal); err != nil {
			log.Fatal(err)
		}
		// } else if cal.Operator == "sub" {
		// 	cal.Result = cal.LeftNum - cal.RightNum
		// 	t, err := template.ParseFiles("template/view.html")
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	if err := t.Execute(w, cal); err != nil {
		// 		log.Fatal(err)
		// 	}
		// } else if cal.Operator == "mult" {
		// 	cal.Result = cal.LeftNum * cal.RightNum
		// 	t, err := template.ParseFiles("template/view.html")
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	if err := t.Execute(w, cal); err != nil {
		// 		log.Fatal(err)
		// 	}
		// } else {
		// 	cal.Result = cal.LeftNum / cal.RightNum
		// 	t, err := template.ParseFiles("template/view.html")
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	if err := t.Execute(w, cal); err != nil {
		// 		log.Fatal(err)
		// 	}
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/calc", calcHandler)
	http.ListenAndServe(":8080", nil)
}
