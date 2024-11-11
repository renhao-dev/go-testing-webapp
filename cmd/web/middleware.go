package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

type contextKey string

const contextIPKey contextKey = "user_ip"

func (app *application) ipFromContext(ctx context.Context) string {
	return ctx.Value(contextIPKey).(string)
}

func (app *application) AddIPToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ip, _ := getIP(r)

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextIPKey, ip)))
		},
	)
}

func getIP(r *http.Request) (string, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	valid := true

	binIP := net.ParseIP(ip)
	if binIP == nil {
		valid = false
		//return "", fmt.Errorf("Invalid ip format")
	}

	forward := r.Header.Get("X-Forwarded-For")
	if len(forward) > 0 {
		binIP = net.ParseIP(forward)
		ip = forward
		valid = binIP != nil
	}

	if !valid {
		ip = "invalid"
		err = fmt.Errorf("Not found valid IP record")
	}

	return ip, err
}
