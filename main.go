package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/aboutme", aboutMeHandler)
	http.HandleFunc("/projects", projectsHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/error", errHandler)

	fs := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources/", fs))

	fmt.Println("Starting server at - http://localhost:8080/")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("\rServer just straight up crashed yo...")
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, r); err != nil {
		panic(err)
	}

}

func aboutMeHandler(w http.ResponseWriter, r *http.Request) {

}

func projectsHandler(w http.ResponseWriter, r *http.Request) {

}

func contactHandler(w http.ResponseWriter, r *http.Request) {

}

func errHandler(w http.ResponseWriter, r *http.Request) {

}
