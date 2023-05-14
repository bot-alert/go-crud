package internal

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type Claims struct {
	Email      string `json:"email"`
	Role       string `json:"role"`
	Identifier int64  `json:"identifier"`
	jwt.RegisteredClaims
}

var JwtKey = []byte("my_secret_key_ashish_123_password_key")

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		token = token[7:]
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if time.Until(claims.ExpiresAt.Time) <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "user", *claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
