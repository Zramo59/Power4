package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("index.html"))
var jeux = NewPower()

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Get the selected column from the form
		col := r.FormValue("col")
		if col != "" {
			var colIdx int
			_, err := fmt.Sscanf(col, "%d", &colIdx)
			if err == nil && colIdx >= 1 && colIdx <= 7 {
				// Play the move (columns are 1-based in UI)
				// Find the lowest empty row in the column
				for i := 5; i >= 0; i-- {
					if jeux.Table[i][colIdx-1] == "" {
						jeux.Table[i][colIdx-1] = jeux.currentPlayer
						jeux.switchPlayer()
						break
					}
				}
			}
		}
	}
	err := tmpl.Execute(w, jeux)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", HandleFunc)
	fmt.Println("Le serveur démarre sur http://localhost:8080")
	// p := NewPower()
	// Starter(p.Player)
	// p.PrintBoard()
	// for {
	// 	if Play(p) {
	// 		if p.Win() {
	// 			break
	// 		} else if p.IsDraw() {
	// 			fmt.Println("Match nul!")
	// 			break
	// 		}
	// 	}
	// }
	http.ListenAndServe(":8080", nil)
}
