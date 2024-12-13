package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var GlobalDatabase *sql.DB

func ConfigDatabase() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred on .env file please check")
	}

	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	localDatabase, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There is a error while connecting to the database", errSql)
		panic(errSql)
	} else {
		GlobalDatabase = localDatabase
		fmt.Println("Successfully connected to database!")
	}
}
