package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	//define port command line flag
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse() //parse the flags so they can be used
	mux := http.NewServeMux()
	urlHandlerMap := map[string]http.HandlerFunc{homeURL: home, showSnippetURL: showSnippet, createSnippetURL: createSnippet}
	for url, handler := range urlHandlerMap {
		mux.HandleFunc(url, handler)
	}
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle(staticURL, http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on port %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
