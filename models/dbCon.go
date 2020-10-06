package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func CreateDB() {
	db := DBConnect()
	db.Migrator().CreateTable(&User{})
	db.Migrator().CreateTable(&Todo{})
}
