package middleware

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func AuthenticatedRedirector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, _ := jwtauth.FromContext(r.Context())

		if token != nil && jwt.Validate(token) == nil {
			http.Redirect(w, r, "/", http.StatusFound)
		}

		next.ServeHTTP(w, r)
	})
}

func UnAuthenticatedRedirector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

		if token == nil || jwt.Validate(token) != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

		next.ServeHTTP(w, r)
	})
}
