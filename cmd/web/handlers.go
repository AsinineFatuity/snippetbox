package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const homeURL string = "/"
const showSnippetURL string = "/snippet"
const createSnippetURL string = "/snippet/create"
const staticURL string = "/static/"

func home(w http.ResponseWriter, r *http.Request) {
	// enforce the home page to be fixed path instead of subtree path
	if r.URL.Path != homeURL {
		http.NotFound(w, r)
		return
	}
	filesToParse := []string{"./ui/html/home.page.tmpl", "./ui/html/base.layout.tmpl", "./ui/html/footer.partial.tmpl"}
	ts, err := template.ParseFiles(filesToParse...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	snippetId, error := strconv.Atoi(r.URL.Query().Get("id"))
	if error != nil || snippetId < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Displaying snippet with ID %d...", snippetId)
}
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		//let client know allowed methods
		w.Header().Set("Allow", http.MethodPost)
		//let the client know that the method is not allowed
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
