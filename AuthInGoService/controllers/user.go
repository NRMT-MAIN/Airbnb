package controllers

import (
	"AuthInGo/dtos"
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
	user , err := uc.userService.GetUserById(userId)

	if err != nil {
		fmt.Println("Error in fetching the user" , err)
		utils.WriteErrorJsonResponse(w , "User Fetching Error" , http.StatusInternalServerError , err)
		return
	}
	
	utils.WriteSuccessJsonResponse(w , "User Found" , http.StatusFound , user)
}

func (uc *UserController) Create(w http.ResponseWriter , r *http.Request) {
	fmt.Println("Register User called in User Controller.")

	payload := r.Context().Value("validatedPayload").(dtos.CreateUserRequest)
	
	uc.userService.CreateUser(&payload)
	utils.WriteSuccessJsonResponse(w , "User Sigup Successfull" , http.StatusAccepted , "")
}

func (uc *UserController) LoginUser(w http.ResponseWriter , r *http.Request){
	fmt.Println("Login User called in User Controller.")
	 
	payload := r.Context().Value("validatedPayload").(dtos.LoginUserRequest)

	jwtToken , err := uc.userService.LoginUser(&payload)
	
	if err != nil {
		utils.WriteErrorJsonResponse(w , "Wrong Credentials" , http.StatusNotAcceptable , err)
		return
	}

	utils.WriteSuccessJsonResponse(w , "User Login Successfull" , http.StatusAccepted , jwtToken)
}