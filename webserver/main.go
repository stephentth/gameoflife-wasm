package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err) // TODO:
	}

	err = template.Execute(w, nil)
	if err != nil {
		panic(err) // TODO:
	}
}

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	staticDir := filepath.Join(exPath, "static")
	// log.Println(staticDir)
	// return

	// fs := http.FileServer(http.Dir("static"))
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveIndex)

	log.Println("Server start!")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Server error:", err)
	}
}
