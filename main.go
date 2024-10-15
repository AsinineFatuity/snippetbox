package main

import (
	"log"
	"net/http"
)

const homeURL string = "/"
const showSnippetURL string = "/snippet"
const createSnippetURL string = "/snippet/create"

func home(w http.ResponseWriter, r *http.Request) {
	// enforce the home page to be fixed path instead of subtree path
	if r.URL.Path != homeURL {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Home page..."))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	urlHandlerMap := map[string]http.HandlerFunc{homeURL: home, showSnippetURL: showSnippet, createSnippetURL: createSnippet}
	for url, handler := range urlHandlerMap {
		mux.HandleFunc(url, handler)
	}
	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
