package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Server runs at http://127.0.0.1:8888, change index.html")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, _ := template.ParseFiles("index.html")
		temp.Execute(w, nil)
	})

	http.ListenAndServe(":8888", nil)
}
