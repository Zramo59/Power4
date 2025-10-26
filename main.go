package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func main() {
	// Créer une nouvelle instance du jeu
	power := NewPower()

	// Parser le template HTML
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Passer la structure Power et Cols au template
		data := struct {
			Power *Power
			Cols  []int
		}{
			Power: power,
			Cols:  []int{0, 1, 2, 3, 4, 5, 6},
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			log.Printf("Erreur lors de l'exécution du template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/difficulty", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Changer la difficulté du jeu
			Difficulty(power)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			colStr := r.FormValue("col")
			// Jouer dans la colonne spécifiée
			col, err := strconv.Atoi(colStr)
			if err == nil {
				Play(power, col)
				power.Win()
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.Handle("/reset", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		power = NewPower()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))

	// Démarrer le serveur web

	log.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
