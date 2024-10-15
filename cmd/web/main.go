package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	urlHandlerMap := map[string]http.HandlerFunc{homeURL: home, showSnippetURL: showSnippet, createSnippetURL: createSnippet}
	for url, handler := range urlHandlerMap {
		mux.HandleFunc(url, handler)
	}
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle(staticURL, http.StripPrefix("/static", fileServer))

	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
