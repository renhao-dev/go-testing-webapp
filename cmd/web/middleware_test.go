package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_AddIPToContext(t *testing.T) {
	tests := []struct {
		name          string
		headerName    string
		headerVal     string
		addr          string
		bEmptyAddress bool
	}{
		{"Natural address", "", "", "", false},
		{"Empty address", "", "", "", true},
		{"Proxy address", "X-Forwarded-For", "192.0.0.10", "192.0.0.1", false},
		{"Invalid custom address", "", "", "192.0.0.3", false},
		{"Custom address", "", "", "192.0.0.3:80", false},
		{"Spoofed address", "", "", "spoof:spoof", false},
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipVal := r.Context().Value(contextIPKey)
		if ipVal == nil {
			t.Error("Value not present")
		}

		ip, ok := ipVal.(string)

		if !ok {
			t.Error("Invalid value in context")
		}
		t.Log(ip)
	})

	for _, test := range tests {
		app := application{}

		req := httptest.NewRequest("GET", "http://testing", nil)
		if len(test.addr) > 0 {
			req.RemoteAddr = test.addr
		}
		if len(test.headerName) > 0 {
			req.Header.Add(test.headerName, test.headerVal)
		}
		if test.bEmptyAddress {
			req.RemoteAddr = ""
		}
		handler := app.AddIPToContext(testHandler)

		t.Logf("%s : ", test.name)

		handler.ServeHTTP(httptest.NewRecorder(), req)
	}
}
