package main

import (
	"beaverops/internal/views"
	"log"
	"net/http"
)

func handleShowCase(w http.ResponseWriter, r *http.Request) {
	views.Showcase().Render(r.Context(), w)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("GET /", handleShowCase)

	log.Println("[+]  BeaverOPS website running at http://localhost:1337/")
	http.ListenAndServe(":1337", mux)
}
