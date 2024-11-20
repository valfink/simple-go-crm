package middleware

import (
	"log/slog"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Incoming request", "Method", r.Method, "URL", r.URL)

		next.ServeHTTP(w, r)
	})
}
