package middlewares

import (
	env "AuthInGo/config/env"
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeaders := r.Header.Get("Authorization")

		if authHeaders == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeaders , "Bearer ") {
			http.Error(w, "Authorization header must start with Bearer", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeaders , "Bearer ")

		if token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}
		parsedToken , err := jwt.ParseWithClaims(token , &claims , func(t *jwt.Token) (interface{}, error) {
			return []byte(env.GetString("JWT_SECRET" , "SECRET")) , nil
		})

		 if err != nil || !parsedToken.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

		userId , okId := claims["id"].(float64)
		email , okEmail := claims["email"].(string)

		if !okId || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context() , "userId" , strconv.Itoa(int(userId)))
		ctx = context.WithValue(ctx , "email" , email)
		next.ServeHTTP(w , r.WithContext(ctx))
	})
}