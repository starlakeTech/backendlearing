package main

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func initDB() {
	//connStr := "user=your_user password=your_password dbname=your_db sslmode=disable"
	connStr := "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable" //ssl协议用不了，123456是密码
	var err error
	db, err = sql.Open("postgres", connStr) //开postgres数据库的connstr地址
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Database connected successfully!")
}

func main() {
	initDB()
}
