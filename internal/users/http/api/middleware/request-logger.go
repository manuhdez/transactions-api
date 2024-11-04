package middleware

import (
	"log"
	"net/http"
)

var RequestLogger = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.Proto)
		next.ServeHTTP(w, r)
	})
}
