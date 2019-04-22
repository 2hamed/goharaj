package main

import (
	"net/http"
	"fmt"
	"context"
)

func Apply(h http.Handler) http.Handler {

	h = JwtAuthMiddleware(h)

	return h
}

func JwtAuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("JWT is Here")
		c := context.WithValue(r.Context(), "saas", "asas")
		handler.ServeHTTP(w, r.WithContext(c))
	})
}
