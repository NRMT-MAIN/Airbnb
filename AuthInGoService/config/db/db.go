package config

import (
	config "AuthInGo/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(){
	var err error
	DB , err = setupDB()

	if err != nil {
		fmt.Println("Error setting up DB connection" , err)
		return
	}
}

func setupDB() (*sql.DB , error){
	cfg := mysql.NewConfig()

	cfg.User = config.GetString("DB_USER" , "root")
	cfg.Passwd = config.GetString("DB_PASSWORD" , "root")
	cfg.Net = config.GetString("DB_TCP" , "tcp")
	cfg.Addr = config.GetString("DB_ADDR" , "127.0.0.1:3306")
	cfg.DBName = config.GetString("DB_NAME" , "AuthDB")

	fmt.Println("Connecting to database:", cfg.DBName, cfg.FormatDSN())
	//from this config it convert to DSN string
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, err
	}

	fmt.Println(("Trying to connect to database..."))
	//verify database connection
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error pinging database:", pingErr)
		return nil, pingErr
	}
	fmt.Println("Connected to database successfully:", cfg.DBName)

	return db, nil
}