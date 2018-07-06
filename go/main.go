package main

import (
	"html/template"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	t, _ := template.ParseFiles("views/views.html")
	t.Execute(w, hostname)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
