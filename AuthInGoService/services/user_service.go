package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dtos"
	"AuthInGo/models"
	"AuthInGo/utils"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById(id string) (*models.User , error)
	GetUserByEmail(email string) (*models.User , error)
	CreateUser(payload *dtos.CreateUserRequest) (*models.User , error)
	GetAllUser() ([]*models.User , error)
	DeleteByUserId(id string) error
	LoginUser(payload *dtos.LoginUserRequest) (string , error)
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (us *UserServiceImpl) DeleteByUserId(id string) error {
	fmt.Println("Deleting user in user service.")
	err := us.userRepository.DeleteById(id)
	if err != nil {
		fmt.Println("Error in deleting user :" , err)
		return err
	}
	return nil
}

func (us *UserServiceImpl) GetAllUser() ([]*models.User , error) {
	fmt.Println("Getting All User in user service.")
	user , err := us.userRepository.GetAll()
	if err != nil {
		fmt.Println("Error getting all user :" , err)
		return nil , err
	}
	return user , nil
}

func (us *UserServiceImpl) CreateUser(payload *dtos.CreateUserRequest) (*models.User , error) {
	fmt.Println("Creating User in user service.")
	hashPassword , err := utils.HashPassword(payload.Password)
	
	if err != nil {
		fmt.Println("Error in getting hashed password")
	}
	
	user , err := us.userRepository.Create(payload.Username , payload.Email , hashPassword)

	if err != nil {
		fmt.Println("Error creating user :" , err)
		return nil , err
	}
	return user , nil
}

func (us *UserServiceImpl) GetUserByEmail(email string) (*models.User , error) {
	fmt.Println("Fetching User in user service.")
	user , err := us.userRepository.GetByEmail(email)

	if err != nil {
		fmt.Println("Error fetching user :" , err)
		return nil , err
	}
	return user , nil
}

func (us *UserServiceImpl) GetUserById(id string) (*models.User , error) {
	fmt.Println("Fetching User in user service.")
	user , err := us.userRepository.GetById(id)

	if err != nil {
		fmt.Println("Error fetching user :" , err)
		return nil , err
	}
	return user , nil
}

func (us *UserServiceImpl) LoginUser(payload *dtos.LoginUserRequest) (string , error) {
	fmt.Println("Login User in user service")

	user , err := us.userRepository.GetByEmail(payload.Email)
	if err != nil {
		fmt.Println("Error in Logging the user" , err)
		return "" , err
	}
	fmt.Println(user.Password , user.Id)

	isPasswordValid := utils.CheckHashPassword(payload.Password ,user.Password)

	if !isPasswordValid {
		fmt.Println("Password valid :" , isPasswordValid)
		return "" , errors.ErrUnsupported
	}

	key := []byte(env.GetString("JWT_SECRET" , "TOKEN"))

	jwtPayload := jwt.MapClaims{
		"email" : user.Email , 
		"id" : user.Id ,
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , jwtPayload)

	tokenString , err := token.SignedString(key)
	if err != nil{
		fmt.Println("Error creating the token")
		return "" , err
	}
	fmt.Println("Token :" , tokenString)
	return tokenString , nil
}