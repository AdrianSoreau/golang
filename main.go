package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"estiam-main/dictionary"
	"estiam-main/middleware"
	"github.com/gorilla/mux"
)

type Word struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

func main() {
	d := dictionary.New("dictionary.json")
	router := mux.NewRouter()

	router.Use(middleware.AuthMiddleware)
	router.Use(middleware.LoggingMiddleware)

	router.HandleFunc("/word", addWord(d)).Methods("POST")
	router.HandleFunc("/word/{word}", getDefinition(d)).Methods("GET")
	router.HandleFunc("/word/{word}", deleteWord(d)).Methods("DELETE")

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}

func addWord(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var word Word
		err := json.NewDecoder(r.Body).Decode(&word)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if !validateWord(word) {
			http.Error(w, "Invalid word or definition", http.StatusBadRequest)
			return
		}

		err = d.Add(word.Word, word.Definition)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to add word: %v", err), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func getDefinition(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		entry, err := d.Get(vars["word"])
		if err != nil {
			http.Error(w, "Word does not exist", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(entry)
	}
}

func deleteWord(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		err := d.Remove(vars["word"])
		if err != nil {
			http.Error(w, "Failed to remove word", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func validateWord(word Word) bool {
	// Ajoutez ici vos règles de validation, par exemple :
	return len(word.Word) > 0 && len(word.Definition) > 0
}
