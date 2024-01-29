package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Car struct {
	Model  string
	Manual bool
}
type PageData struct {
	Title string
	Cars  []Car
}

func car(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "list of cars",
		Cars: []Car{
			{Model: "Volvo v70", Manual: true},
			{Model: "Toyota Prius", Manual: false},
			{Model: "Audi A6 avant", Manual: false},
		},
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	tmpl = template.Must(template.ParseFiles("/Users/wen/GolandProjects/haorui/static/templates/index.gohtml"))
	mux.HandleFunc("/car", car)
	log.Fatal(http.ListenAndServe(":9091", mux))
}
