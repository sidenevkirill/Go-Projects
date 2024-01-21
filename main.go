package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/home.html")
	tmpl.Execute(w, nil)
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/home", sayhello)
	http.HandleFunc("/main", hello)
	http.HandleFunc("/headers", headers)

	//	http.ListenAndServe("141.8.192.193:8080", nil)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
