package middlewares

import (
	"AuthInGo/utils"
	"errors"
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimiter(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(5 , 5)
	return http.HandlerFunc(func(w http.ResponseWriter , r *http.Request) {
		if !limiter.Allow(){
			utils.WriteErrorJsonResponse(w , "Too many Request" , http.StatusTooManyRequests , errors.New("Request Reached"))
			return 
		}
		next.ServeHTTP(w , r)
	})
}