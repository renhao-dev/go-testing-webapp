package main

import (
	"html/template"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var pathToTemplates = "./templates/"

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.AddIPToContext)

	mux.Get("/", app.Home)
	mux.Post("/login", app.Login)
	mux.Get("/showip", app.ShowIP)

	return mux
}

type TemplateData struct {
	IP   string
	data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	data.IP, _ = app.ipFromContext(r.Context())

	parsedTempl, err := template.ParseFiles(path.Join(pathToTemplates, t))

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
