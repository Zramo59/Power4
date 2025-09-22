// package main

// import (
// 	"html/template"
// 	"net/http"
// )

// var tmpl = template.Must(template.ParseFiles("index.html"))

// func handler(w http.ResponseWriter, r *http.Request) {
// 	data := struct{ Titre string }{Titre: "Bonjour, Monde!"}
// 	tmpl.Execute(w, data)
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":8080", nil)
// 	NewPower()
// }
