package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string) (string , error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword) , bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password" ,err)
		return "" , err
	}
	return string(hashedPassword) , nil
}

func CheckHashPassword(plainPassword string , hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword) , []byte(plainPassword))
	return err == nil
}