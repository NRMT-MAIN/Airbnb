package middlewares

import (
	"AuthInGo/utils"
	"context"
	"fmt"
	"net/http"
)

func ValidateRequestBody[T any](next http.HandlerFunc)http.HandlerFunc {
	return func(w http.ResponseWriter , r *http.Request){
		var payload T
		if jsonErr := utils.ReadJsonBody(r , &payload) ; jsonErr != nil {
			utils.WriteErrorJsonResponse(w , "JSON Reading Error" , http.StatusInternalServerError , jsonErr)
			return 
		}
		fmt.Println(payload)
		if validateErr := utils.Validator.Struct(payload) ; validateErr != nil {
			utils.WriteErrorJsonResponse(w , "Validation Failed" , http.StatusNotAcceptable , validateErr)
			return 
		}
		
		ctx := context.WithValue(r.Context() , "validatedPayload" , payload)
		next.ServeHTTP(w ,r.WithContext(ctx))
	}
}

