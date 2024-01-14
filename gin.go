package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func mysqlPoolConnection() (*sql.DB, error) {
	// db, err := sql.Open("mysql", "root:Aeiou@1234/BookSchema")
    db, err := sql.Open("mysql", "root:Aeiou@1234@tcp(127.0.0.1:3306)/BookSchema")

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// I will refer the below link for the above db methods
	// https://github.com/go-sql-driver/mysql#important-settings

	return db, nil
}

func main() {
	db, err := mysqlPoolConnection()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}

	fmt.Println("Successfully connected to the database")
}
