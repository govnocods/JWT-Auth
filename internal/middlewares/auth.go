package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/govnocods/JWT-Authorization/internal/auth"
	"github.com/govnocods/JWT-Authorization/internal/db"
	"github.com/govnocods/JWT-Authorization/utils"
)

type Middleware struct {
	DB *db.DataBase
}

func NewMiddleware(database *db.DataBase) *Middleware {
	return &Middleware{
		DB: database,
	}
}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Unauthorization", http.StatusUnauthorized)
				return
			}
			c = &http.Cookie{Value: strings.TrimPrefix(authHeader, "Bearer ")}
		}

		tokenStr := c.Value
		claims, err := auth.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		user, _ := m.DB.GetUser(claims.Username)

		ctx := context.WithValue(r.Context(), utils.UserCtxKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
