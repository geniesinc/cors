package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/axiomzen/cors/v2"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	// Use default options
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8080", handler)
}
