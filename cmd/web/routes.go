package main

import (
	"fmt"
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
	mux.Use(app.Session.LoadAndSave)

	mux.Get("/", app.Home)
	mux.Post("/login", app.Login)
	mux.Get("/showip", app.ShowIP)

	return mux
}

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	data.IP, _ = app.ipFromContext(r.Context())
	fmt.Printf("Session val: %s \n", data.Data["test"].(string))

	parsedTempl, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "layout.base.gohtml"))

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
