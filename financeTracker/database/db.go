package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"

)
var DB *sql.DB
func InitDB(){
	if err:=godotenv.Load();err!=nil{
		log.Fatal("Error loading .env file")
	}
	dbURL:=os.Getenv("DataBase_url")
	var err error
	DB,err=sql.Open("postgres",dbURL)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
    DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	if err:=DB.Ping();err!=nil{
		log.Fatal("Dtabase connection is not active")
	}

	createTables()
	fmt.Println("Connected to the database successfully ")

}
func createTables(){
	usersTable := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`

	transactionsTable := `CREATE TABLE IF NOT EXISTS transactions (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id) ON DELETE CASCADE,
		amount DECIMAL(10,2) NOT NULL,
		category VARCHAR(255) NOT NULL,
		description TEXT,
		date TIMESTAMP DEFAULT now()
	);`

	_,err:=DB.Exec(usersTable)
	if err!=nil{
		log.Fatal("error creating using table")
		return
	}
	_,err=DB.Exec(transactionsTable)
	if err!=nil{
		log.Fatal("error creating transaction table")
		return
	}
  

}
