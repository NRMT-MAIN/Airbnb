package middlewares

import (
	config "AuthInGo/config/db"
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
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

func RequireAllRoles(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := r.Context().Value("userId").(string)

			userIdInt , _ := strconv.Atoi(userId)

			urr := db.NewUserRoleRepository(config.DB)

			hasAllRole , err :=  urr.HasAllRoles(int64(userIdInt) , roles)

			if err != nil {
				http.Error(w, "Error checking user roles: "+ err.Error(), http.StatusInternalServerError)
				return
			}

			if !hasAllRole {
				http.Error(w, "Forbidden: You do not have the required roles", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w , r)
		})
	}
}

func RequireAnyRoles(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := r.Context().Value("userId").(string)

			userIdInt , _ := strconv.Atoi(userId)

			urr := db.NewUserRoleRepository(config.DB)

			hasAllRole , err :=  urr.HasAnyRole(int64(userIdInt) , roles)

			if err != nil {
				http.Error(w, "Error checking user roles: "+ err.Error(), http.StatusInternalServerError)
				return
			}

			if !hasAllRole {
				http.Error(w, "Forbidden: You do not have the required roles", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w , r)
		})
	}
}