package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	r.HandleFunc("/", IndexHandler)

	http.ListenAndServe(":7001", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := ""
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	if err := tmpl.ExecuteTemplate(w, "index", data); err != nil {
		log.Fatal(err)
	}
}
