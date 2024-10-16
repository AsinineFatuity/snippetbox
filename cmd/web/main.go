package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal("Error loading .env file")
	}
	//define port command line flag
	defaultPort := os.Getenv("SNIPPETBOX_ADDR")
	addr := flag.String("addr", defaultPort, "HTTP network address")
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
