package db

import (
	"AuthInGo/dtos"
	"AuthInGo/models"
	"database/sql"
	"fmt"
	"time"
)

type RoleRepository interface {
	GetById(id int64) (*models.Role , error)
	GetByName(name string) (*models.Role , error)
	GetAll() ([]*models.Role , error)
	Create(payload *dtos.CreateRoleRequest) (*models.Role , error)
	DeleteById(id int64) error 
	UpdateById(id int64 , payload *dtos.CreateRoleRequest) (*models.Role , error)
}

type RoleRepositoryImpl struct {
	db *sql.DB
}

func NewRoleRepository(_db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db : _db , 
	}
}


func (r *RoleRepositoryImpl) GetById(id int64) (*models.Role , error) {
	query := "SELECT * FROM ROLE WHERE ID = ?"

	row := r.db.QueryRow(query , id)

	role := &models.Role{}
	err := row.Scan(&role.Id , &role.Name , &role.Description , &role.Created_at , &role.Updated_at)

	if err != nil { 
		if err == sql.ErrNoRows {
			fmt.Println("No record Found!")
		} else {
			fmt.Println("Error scanning row")
		}
		return nil , err
	}

	fmt.Println("Role Fetched Successfully")
	return role , nil
}

func (r *RoleRepositoryImpl) GetByName(name string) (*models.Role , error) {
	query := "SELECT * FROM ROLE WHERE NAME = ?;"

	row := r.db.QueryRow(query , name)
	role := &models.Role{}

	err := row.Scan(&role.Id , &role.Name , &role.Description , &role.Created_at , &role.Updated_at)
	
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No records found!")
		} else { 
			fmt.Println("Error scanning the row")
		}
		return nil , err
	}
	return role , nil
}

func (r *RoleRepositoryImpl) GetAll() ([]*models.Role , error) {
	query := "SELECT * FROM ROLE;"

	rows , err := r.db.Query(query)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No records found!")
		} else { 
			fmt.Println("Error scanning the row")
		}
		return nil , err
	}

	var roles []*models.Role
	for rows.Next() {
		role := &models.Role{}
		err := rows.Scan(&role.Id , &role.Name , &role.Description , &role.Created_at , &role.Updated_at)
		
		if err != nil {
			fmt.Println("Error in scannning role.")
			return nil , err
		}
		roles = append(roles, role)
	}
	return roles , nil
}

func (r *RoleRepositoryImpl) Create(payload *dtos.CreateRoleRequest) (*models.Role , error) {
	query := "INSERT INTO ROLE ( NAME , DESCRIPTION ) VALUES ( ? , ?) ;"

	result , err := r.db.Exec(query , payload.Name , payload.Description)

	if err != nil {
		fmt.Println("Error in creating the role!" , err)
		return nil , err
	}

	lastEnterId , lastInsertIdErr := result.LastInsertId()

	if lastInsertIdErr != nil {
		fmt.Println("Error in getting last enteredId" , lastInsertIdErr)
		return nil , lastInsertIdErr
	}

	role := &models.Role{
		Id: lastEnterId,
		Name: payload.Name,
		Description: payload.Description,
		Created_at: time.Now().String(),
		Updated_at: time.Now().String(),
	}

	fmt.Println("Role created successfully!")
	return role , nil 
}

func (r *RoleRepositoryImpl) DeleteById(id int64) error {
	query := "DELETE FROM ROLE WHERE ID = ?;"

	result , err := r.db.Exec(query , id)

	if err != nil { 
		if err == sql.ErrNoRows {
			fmt.Println("No row found to delete")
		} else  {
			fmt.Println("Error in executing the query")
		}
		return err
	}

	rowsAffected , rowsAffectedErr := result.RowsAffected()
	if rowsAffectedErr != nil {
		fmt.Println("Error getting rows affected:", rowsAffectedErr)
		return rowsAffectedErr
	}
	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not deleted")
		return rowsAffectedErr
	}

	fmt.Println("Row deleted succesfully!")
	return nil 
}

func (r *RoleRepositoryImpl) UpdateById(id int64 , payload *dtos.CreateRoleRequest) (*models.Role , error){
	query := "UPDATE role SET name = COALESCE(?, name), description = COALESCE(?, description) WHERE id = ?"
	result , err := r.db.Exec(query, payload.Name , payload.Description , id)

	if err != nil { 
		if err == sql.ErrNoRows {
			fmt.Println("No row found to update")
		} else  {
			fmt.Println("Error in executing the query")
		}
		return nil , err
	}
	rowsAffected , rowsAffectedErr := result.RowsAffected()
	if rowsAffectedErr != nil {
		fmt.Println("Error getting rows affected:", rowsAffectedErr)
		return nil , rowsAffectedErr
	}
	if rowsAffected == 0 {
		fmt.Println("No rows were affected, role not updated")
		return nil , rowsAffectedErr
	}

	fmt.Println("Row updated succesfully!")

	role := &models.Role{
		Id: id,
		Name: payload.Name,
		Description: payload.Description,
		Created_at: "",
		Updated_at: time.Now().String(),
	}
	return role , nil 
}