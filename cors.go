package middleware

import (
	"net/http"
	"strings"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("access-control-allow-credentials", "true")
		w.Header().Set("access-control-allow-methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

		origin := r.Header.Get("origin")
		if len(strings.TrimSpace(origin)) == 0 {
			origin = "*"
		}
		w.Header().Set("access-control-allow-origin", origin)

		ac := r.Header.Get("access-control-request-headers")
		if len(strings.TrimSpace(ac)) == 0 {
			ac = "*"
		}
		w.Header().Set("access-control-allow-headers", ac)
		if r.Method == "OPTIONS" && bypassOnOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}
