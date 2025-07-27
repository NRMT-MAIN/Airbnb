package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load() ; 
}

func GetString(key string , fallback string) string {
	value , ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}
	return value 
}

func GetInt(key string , fallback int) int{
	value , ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	result , err := strconv.Atoi(value)
	
	if err != nil {
		fmt.Println("Error in string to int conversion in env.")
		return fallback ; 
	}

	return result
}

func GetBool(key string , fallback bool) bool{
	value , ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	result , err := strconv.ParseBool(value)
	
	if err != nil {
		fmt.Println("Error in string to int conversion in env.")
		return fallback ; 
	}

	return result
}