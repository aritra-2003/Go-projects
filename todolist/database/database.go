package database

import (
	"database/sql"
	"fmt"
	"log"
	
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "host=localhost user=postgres password=ARga12@@ dbname=todolist port=5433 sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot reach the database:", err)
	}

	fmt.Println("Connected to the database successfully!")
}