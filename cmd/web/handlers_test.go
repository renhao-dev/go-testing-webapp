package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	tests := []struct {
		name   string
		route  string
		status int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/home", http.StatusNotFound},
	}

	pathToTemplates = "./../../templates/"

	routes := app.routes()

	server := httptest.NewTLSServer(routes)

	for _, test := range tests {
		resp, err := server.Client().Get(server.URL + test.route)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != test.status {
			t.Errorf("%s: Expected status code %d, got %d", test.name, test.status, resp.StatusCode)
		}
	}
}
