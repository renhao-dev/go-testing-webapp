package main

import (
	"fmt"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
}

func (app *application) ShowIP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "IP: %s", app.ipFromContext(r.Context()))
}
