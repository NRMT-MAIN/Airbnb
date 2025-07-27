package config

import (
	"database/sql"
	"fmt"
	"reviewservice/config/env"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB , error){
	cfg := mysql.NewConfig()  

	cfg.Addr = env.GetString("DB_ADDR" , "127.0.0.1:3306")
	cfg.User = env.GetString("DB_USER" , "root")
	cfg.Passwd = env.GetString("DB_PASSWORD" , "root")
	cfg.Net = env.GetString("DB_NET" , "tcp")
	cfg.DBName = env.GetString("DB_NAME" , "airbnbReview_db")

	fmt.Println("Connecting to Database:" , cfg.FormatDSN())

	db , err := sql.Open("mysql" , cfg.FormatDSN())

	if err != nil {
		fmt.Println("Error connecting to database")
		return nil , err 
	}

	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println("Error pinging the database" , pingErr)
		return nil , pingErr
	}

	return db , nil 
}