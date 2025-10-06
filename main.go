package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	// Créer une nouvelle instance du jeu
	power := NewPower()

	// Parser le template HTML
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Passer la structure Power au template
		data := struct {
			Power *Power
		}{
			Power: power,
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			log.Printf("Erreur lors de l'exécution du template: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
