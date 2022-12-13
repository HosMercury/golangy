package main

import (
	"log"
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/base.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8010", nil)
}
