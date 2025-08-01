package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
	"strings"
)

type UserRoleRepository interface {
	GetUserRole(userId int64) ([]*models.Role , error)
	AssignRoleToUser(userId int64 , roleId int64) error
	RemoveRoleFromUser(userId int64, roleId int64) error
	GetUserPermissions(userId int64) ([]*models.Permission, error)
	HasPermission(userId int64, permissionName string) (bool, error)
	HasRole(userId int64, roleName string) (bool, error)
	HasAllRoles(userId int64, roleNames []string) (bool, error)
	HasAnyRole(userId int64, roleNames []string) (bool, error)
}

type UserRoleRepositoryImpl struct {
	db *sql.DB
}

func NewUserRoleRepository(_db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImpl{
		db : _db ,
	}
}

func (u *UserRoleRepositoryImpl) GetUserRole(userId int64) ([]*models.Role , error){
	query := `SELECT R.ID , R.NAME , R.DESCRIPTION , R.CREATED_AT , R.UPDATED_AT 
			  FROM USER_ROLE UR
			  INNER JOIN ROLE R ON R.ID = UR.ID 
			  WHERE USER_ID = ?;`
	
	rows , err := u.db.Query(query , userId)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No role found for particular user id" , err)
		} else {
			fmt.Println("Error in quering" , err)
		}
		return nil , err
	}

	defer rows.Close()

	var roles []*models.Role
	
	for rows.Next() {
		role := &models.Role{}
		err := rows.Scan(&role.Id , &role.Name , &role.Description , &role.Created_at , &role.Updated_at)

		if err != nil {
			fmt.Println("Error in scanning the row" , err)
			return nil ,err
		}
		roles = append(roles, role)
	}
	return roles , nil 
}

func (u *UserRoleRepositoryImpl) AssignRoleToUser(userId int64 , roleId int64) error {
	query := `INSERT INTO USER_ROLE (USER_ID , ROLE_ID) VALUES ( ? , ?);`

	_ , err := u.db.Exec(query , userId , roleId)

	if err != nil {
		fmt.Println("Error executing query" , err)
		return err
	}
	return nil 
}

func (u *UserRoleRepositoryImpl) RemoveRoleFromUser(userId int64 , roleId int64) error {
	query := `DELETE FROM USER_ROLE WHERE USER_ID = ? AND ROLE_ID = ?;`

	_ , err := u.db.Exec(query , userId , roleId)
	
	if err != nil {
		fmt.Println("Error executing query" , err)
		return err
	}
	return nil 
}

func (u *UserRoleRepositoryImpl) GetUserPermissions(userId int64) ([]*models.Permission , error){
	query := `SELECT P.NAME , P.DESCRIPTION , P.RESOURCE , P.ACTION , P.CREATED_AT , P.UPDATED_AT
				FROM USER_ROLE UR 
				INNER JOIN ROLE_PERMISSION RP ON UR.ROLE_ID = RP.ROLE_ID
				INNER JOIN PERMISSION P ON RP.PERMISSION_ID = P.ID
				WHERE UR.USER_ID = ?;`
	
	rows , err := u.db.Query(query , userId)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No permission found for particular user id" , err)
		} else {
			fmt.Println("Error in quering" , err)
		}
		return nil , err
	}

	defer rows.Close()

	var permissions []*models.Permission
	
	for rows.Next() {
		permission := &models.Permission{}
		err := rows.Scan(&permission.Id , &permission.Name , &permission.Description , &permission.Resource , &permission.Action , 
		&permission.CreatedAt , &permission.UpdatedAt)

		if err != nil {
			fmt.Println("Error in scanning the row" , err)
			return nil ,err
		}
		permissions = append(permissions, permission)
	}
	return permissions , nil
}

func (u *UserRoleRepositoryImpl) HasPermission(userId int64 , permissionName string) (bool , error) {
	query := `SELECT COUNT(*) > 0
				FROM USER_ROLE UR
				INNER JOIN ROLE_PERMISSION RP ON RP.ROLE_ID = UR.ROLE_ID
				INNER JOIN PERMISSION P ON P.ID = RP.PERMISSION_ID
				WHERE UR.USER_ID = ? AND P.NAME = ?;`
	
	var exists bool 
	err := u.db.QueryRow(query , userId , permissionName).Scan(&exists)

	if err != nil {
		fmt.Println("Given Permission doesn't associate with user")
		return false , err
	}
	return exists , nil 
}

func (u *UserRoleRepositoryImpl) HasRole(userId int64, roleName string) (bool, error) {
	query := `SELECT COUNT(*)
				FROM USER_ROLE UR
				INNER JOIN ROLE R UR.ROLE_ID = R.ID
				WHERE USER_ROLE = ? AND ROLE.NAME = ?;`
	
	var exists bool ; 
	err := u.db.QueryRow(query , userId , roleName).Scan(&exists)
	if err != nil {
		fmt.Println("Given Role doesn't associate with user")
		return false, err
	}
	return exists, nil
}

func (u *UserRoleRepositoryImpl) HasAllRoles(userId int64, roleNames []string) (bool, error) {
	if len(roleNames) == 0 {
		return true , nil 
	}

	placeholders := make([]string , len(roleNames))
	args := make([]interface{} , len(roleNames) + 2)
	args[0] = len(roleNames)
	args[1] = userId

	for i , value := range roleNames {
		placeholders[i] = "?"
		args[i + 2] = value
	}

	query := `SELECT COUNT(*) = ?
				FROM USER_ROLE UR
				INNER JOIN ROLE R ON R.ID = UR.ROLE_ID
				WHERE UR.USER_ID = ? AND R.NAME IN (` + strings.Join(placeholders , ",") + `) ;`
	var exists bool 
	err := u.db.QueryRow(query , args...).Scan(&exists)

	if err != nil {
		fmt.Println("Error in quering" , err)
		return false , err
	}
	return exists , nil
}

func (u *UserRoleRepositoryImpl) HasAnyRole(userId int64 , roleNames []string) (bool, error) {
	if len(roleNames) == 0 {
		return true , nil 
	}
	placeholders := make([]string, len(roleNames))
	args := make([]interface{}, len(roleNames)+1)
	args[0] = userId
		
	for i, name := range roleNames {
		placeholders[i] = "?"
		args[i+1] = name
	}

	query := `SELECT COUNT(*) > 0 
				FROM USER_ROLE UR
				INNER JOIN ROLE R ON R.ID = UR.ROLE_ID
				WHERE UR.USER_ID = ? AND R.NAME IN (` + strings.Join(placeholders, ",") + `)`
	var exists bool 
	err := u.db.QueryRow(query , args...).Scan(&exists)

	if err != nil {
		fmt.Println("Error in quering" , err)
		return false , err
	}
	return exists , nil
}