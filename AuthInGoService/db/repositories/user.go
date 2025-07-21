package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetById(id string) (*models.User , error)
	GetAll() ([]*models.User , error)
	Create(username string , email string , hashedPassword string) (*models.User , error)
	GetByEmail(email string) (*models.User , error)
	DeleteById(id string) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) DeleteById(id string) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := u.db.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting user:", err)
		return err
	}

	rowsAffected, rowErr := result.RowsAffected()
	if rowErr != nil {
		fmt.Println("Error getting rows affected:", rowErr)
		return rowErr
	}
	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not deleted")
		return nil
	}
	fmt.Println("User deleted successfully, rows affected:", rowsAffected)
	return nil
}

func (u *UserRepositoryImpl) GetByEmail(email string) (*models.User , error){
	query := "SELECT * FROM USERS WHERE EMAIL = ?"

	row := u.db.QueryRow(query , email)
	
	user := &models.User{}
	err := row.Scan(&user.Id , &user.Username , &user.Email , &user.Password , &user.CreatedAt , &user.UpdatedAt)

	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("No Records Found!")
			return nil , err
		} else {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
	}

	return user , nil
}

func (u *UserRepositoryImpl) Create(username string , email string , hashedPassword string) (*models.User , error) {
	query := "INSERT INTO USERS (USERNAME , EMAIL , PASSWORD) VALUES ( ? , ? , ?)"

	result , err := u.db.Exec(query , username , email , hashedPassword)
	if err != nil {
		fmt.Println("Error creating the user")
		return nil , err
	}

	lastEnterId , err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error in getting last entered id" , err)
		return nil , err
	}

	user := &models.User{
		Id: lastEnterId,
		Username: username,
		Email: email,
	}

	fmt.Println("User created Succesfully" , user)
	return user , nil
}


func (u *UserRepositoryImpl) GetAll() ([]*models.User , error){
	query := "SELECT * FROM USERS"

	rows ,err := u.db.Query(query)

	if err != nil {
		fmt.Println("Error fetching the rows from DB" , err)
		return nil, err
	}

	defer rows.Close() // Ensure rows are closed after processing

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id , &user.Username , &user.Email , &user.Password , &user.CreatedAt , &user.UpdatedAt)
		if err != nil {
			fmt.Println("Error fetching the user" , err)
			return nil , err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error with rows:", err)
		return nil , err
	}

	return users , nil 
}

func (u *UserRepositoryImpl) GetById(id string)(*models.User , error){
	fmt.Println("Getting in User Repository")

	//sql injections
	query := "SELECT * FROM USERS WHERE ID = ?"
	row := u.db.QueryRow(query , id)

	user := &models.User{}

	err := row.Scan(&user.Id , &user.Username , &user.Email , &user.Password , &user.CreatedAt , &user.UpdatedAt)

	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("No Records Found!")
			return nil , err
		} else {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
	}
	fmt.Println("Uer fetched Successfully :" , user)
	return user, nil
}