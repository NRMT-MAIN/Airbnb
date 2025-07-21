package controllers

import (
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		userService: _userService,
	}
}

func (uc *UserController) GetUserById(w http.ResponseWriter , r *http.Request){
	fmt.Println("Register User called in User Controller.")
	userId := r.URL.Query().Get("id")
	uc.userService.GetUserById(userId)
	w.Write([]byte("User Registration Endpoint."))
}

func (uc *UserController) Create(w http.ResponseWriter , r *http.Request) {
	fmt.Println("Register User called in User Controller.")
	hashPassword , err := utils.HashPassword("nirmit")
	if err != nil {
		fmt.Println("Error in getting hashed password")
	}
	uc.userService.CreateUser("u_demo1" , "u1@demo.com" , hashPassword)
	w.Write([]byte("User Registration Endpoint."))
}

func (uc *UserController) LoginUser(w http.ResponseWriter , r *http.Request){
	fmt.Println("Login User called in User Controller.")
	jwtToken , err := uc.userService.LoginUser()
	
	if err != nil {
		w.Write([]byte("Something went wrong!"))
	}

	response := map[string]any{
		"status" : "sucess" ,
		"token" : jwtToken ,
		"error" : nil , 
	}

	utils.WriteJsonResponse(w , http.StatusAccepted , response)
}