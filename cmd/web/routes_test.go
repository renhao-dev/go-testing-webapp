package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_application_routes(t *testing.T) {
	registered := []struct {
		routePath string
		method    string
	}{
		{"/", "GET"},
	}
	mux := app.routes()
	routes := mux.(chi.Routes)

	for _, route := range registered {
		if !routeExists(route.routePath, route.method, routes) {
			t.Errorf("Route %s with method %s is expected to be registered but is not", route.routePath, route.method)
		}
	}
}

func routeExists(routePath string, routeMethod string, routes chi.Routes) bool {
	found := false
	chi.Walk(routes, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == routePath && method == routeMethod {
			found = true
		}
		return nil
	})

	return found
}
