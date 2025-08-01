package middlewares

import (
	"AuthInGo/dtos"
	"AuthInGo/utils"
	"context"
	"fmt"
	"net/http"
)

func ValidateRequestBody[T any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter , r *http.Request){
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
	})
}


func UserLoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dtos.LoginUserRequest

		// Read and decode the JSON body into the payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteErrorJsonResponse(w, "Invalid request body", http.StatusBadRequest, err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteErrorJsonResponse(w, "Validation failed", http.StatusBadRequest, err)
			return
		}

		fmt.Println("Payload received for login:", payload)

		ctx := context.WithValue(r.Context(), "payload", payload) // Create a new context with the payload

		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}

func UserCreateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dtos.CreateUserRequest

		// Read and decode the JSON body into the payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteErrorJsonResponse(w, "Invalid request body", http.StatusBadRequest, err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteErrorJsonResponse(w, "Validation failed", http.StatusBadRequest, err)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}

func CreateRoleRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dtos.CreateRoleRequestDTO

		// Read and decode the JSON body into the payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteErrorJsonResponse(w, "Invalid request body", http.StatusBadRequest, err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteErrorJsonResponse(w, "Validation failed", http.StatusBadRequest, err)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UpdateRoleRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dtos.UpdateRoleRequestDTO

		// Read and decode the JSON body into the payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteErrorJsonResponse(w, "Invalid request body", http.StatusBadRequest, err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteErrorJsonResponse(w, "Validation failed", http.StatusBadRequest, err)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AssignPermissionRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dtos.AssignPermissionRequestDTO

		// Read and decode the JSON body into the payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteErrorJsonResponse(w, "Invalid request body", http.StatusBadRequest, err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteErrorJsonResponse(w, "Validation failed", http.StatusBadRequest, err)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RemovePermissionRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dtos.RemovePermissionRequestDTO

		// Read and decode the JSON body into the payload
		if err := utils.ReadJsonBody(r, &payload); err != nil {
			utils.WriteErrorJsonResponse(w, "Invalid request body", http.StatusBadRequest, err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			utils.WriteErrorJsonResponse(w, "Validation failed", http.StatusBadRequest, err)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}