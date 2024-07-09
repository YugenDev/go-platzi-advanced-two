package middleware

import (
	"net/http"
	"strings"

	"github.com/YugenDev/go-platzi-advanced-two/models"
	"github.com/YugenDev/go-platzi-advanced-two/server"
	"github.com/golang-jwt/jwt/v4"
)

var (
	NO_AUTH_NEEDED = []string{
		"login",
		"signup",
	}
)

func ShouldCheckToken(route string) bool {
	for _, path := range NO_AUTH_NEEDED {
		if strings.Contains(route, path) {
			return false
		}
	}
	return true
}

func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !ShouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			_, err := jwt.ParseWithClaims(tokenString, models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecretKey), nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
