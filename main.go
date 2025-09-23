package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("index.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title  string
		grille [6][7]string
	}{
		Title:  "Bienvenue sur mon puissance 4",
		grille: [6][7]string{},
	}
	tmpl.Execute(w, data)
	// Afficher un message de bienvenue
	fmt.Fprintln(w, "Bienvenue sur mon puissance 4")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	// Initialize the game (this is just an example, actual game logic would be more complex)
	// powerGame := NewPower()
}
