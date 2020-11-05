package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DBConnect() *sqlx.DB {
	env := os.Getenv("GIN_MODE")

	var host, port, user, password, dbname string
	if env == "release" {
		host = os.Getenv("dbHost")
		port = os.Getenv("dbPort")
		user = os.Getenv("dbUser")
		password = os.Getenv("dbPassword")
		dbname = os.Getenv("dbname")
	} else {
		host = "localhost"
		port = "3306"
		user = "root"
		password = "admin"
		dbname = "gin_demo"
	}

	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)
	db, err := sqlx.Connect("mysql", dbConfig)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return db
}
