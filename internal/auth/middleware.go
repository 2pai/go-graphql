package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/2pai/go-graphql/internal/pkg/jwt"
	"github.com/2pai/go-graphql/internal/users"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"User"}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("token")

			if err != nil || c == nil {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := c.Value
			username, err := jwt.ParseToken(tokenStr)

			if err != nil {
				http.Error(w, "Invalid Token", http.StatusForbidden)
				return
			}

			user := users.User{Username: username}
			id, err := users.GetUserIdByUsername(username)

			user.ID = strconv.Itoa(id)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, user)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
