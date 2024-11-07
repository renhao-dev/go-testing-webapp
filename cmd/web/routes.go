package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", app.Home)

	return mux
}

type TemplateData struct {
	IP   string
	data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	parsedTempl, err := template.ParseFiles("./templates/" + t)

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	err = parsedTempl.Execute(w, data)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return err
	}

	return nil
}
