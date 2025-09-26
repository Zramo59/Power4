package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("index.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Serveur démarré sur http://localhost:8080")
	//
	http.ListenAndServe(":8080", nil)
	p := NewPower()
	Starter(p.Player)
	p.PrintBoard()
	for {
		if Play(p) {
			if p.Win() {
				break
			} else if p.IsDraw() {
				fmt.Println("Match nul!")
				break
			}
		}
	}
}
