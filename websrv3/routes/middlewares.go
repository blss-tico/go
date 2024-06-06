package routes

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("\n Method: [%v] - Route: [%v] - Remote Address: [%v]", r.Method, r.RequestURI, r.RemoteAddr)
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		},
	)
}

func SetHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		},
	)
}
