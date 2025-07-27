package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init(){
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func WriteJsonResponse(w http.ResponseWriter , status int , data any) error {
	w.Header().Set("Content-Type" , "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func WriteSuccessJsonResponse(w http.ResponseWriter , message string , status int , data any) error {
	response := map[string]any{}

	response["message"] = message 
	response["status"] = "SUCCESS"
	response["data"] = data
	
	return WriteJsonResponse(w , status , response)
}

func WriteErrorJsonResponse(w http.ResponseWriter , message string , status int , err error) error {
	response := map[string]any{}

	response["message"] = message 
	response["status"] = "ERROR"
	response["Error"] = err.Error()
	
	return WriteJsonResponse(w , status , response)
}

func ReadJsonBody(r *http.Request , result any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(result)
}