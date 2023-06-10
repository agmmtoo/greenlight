package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status:\tavailable")
	fmt.Fprintf(w, "environment:\t%s\n", app.config.env)
	fmt.Fprintf(w, "version:\t%s\n", version)
}
