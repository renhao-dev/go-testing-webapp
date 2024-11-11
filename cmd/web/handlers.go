package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("pwd")
	log.Printf("email: %s; password: %s", email, password)

	fmt.Fprint(w, email)
}

func (app *application) ShowIP(w http.ResponseWriter, r *http.Request) {
	ip, _ := app.ipFromContext(r.Context())
	fmt.Fprintf(w, "IP: %s", ip)
}
