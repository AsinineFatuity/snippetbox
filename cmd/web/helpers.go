package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// writes traceback for error message and returns a 500 internal server error
func (app *application) serverError(w http.ResponseWriter, err error) {
	// get the stack trace for the current go routine and append to the log message
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// sends a client error with the status code
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// sends a not found error to the client
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
