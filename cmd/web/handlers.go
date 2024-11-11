package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	td := make(map[string]any)

	if app.Session.Exists(r.Context(), "test") {
		td["test"] = app.Session.GetString(r.Context(), "test")
		log.Printf("Session exists! Got %s \n", td["test"])
	} else {
		td["test"] = fmt.Sprintf("Hit page at %s", time.Now().String())
		app.Session.Put(r.Context(), "test", td["test"])
	}
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{Data: td})
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	form := NewForm(r.PostForm)
	form.Required("email", "pwd")

	if !form.Valid() {
		fmt.Fprint(w, "Validation failed")
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
