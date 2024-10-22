package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	envError := godotenv.Load()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	if envError != nil {
		errorLog.Fatal("Error loading .env file")
	}
	//define port command line flag
	defaultPort := os.Getenv("SNIPPETBOX_ADDR")
	addr := flag.String("addr", defaultPort, "HTTP network address")
	flag.Parse() //parse the flags so they can be used
	mux := http.NewServeMux()
	// define new instance of app containing the dependecies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	urlHandlerMap := map[string]http.HandlerFunc{homeURL: app.home, showSnippetURL: app.showSnippet, createSnippetURL: app.createSnippet}
	for url, handler := range urlHandlerMap {
		mux.HandleFunc(url, handler)
	}
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle(staticURL, http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Starting server on port %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
