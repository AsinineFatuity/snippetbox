package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	urlHandlerMap := map[string]http.HandlerFunc{homeURL: app.home, showSnippetURL: app.showSnippet, createSnippetURL: app.createSnippet}
	for url, handler := range urlHandlerMap {
		mux.HandleFunc(url, handler)
	}
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle(staticURL, http.StripPrefix("/static", fileServer))
	return mux
}
