package db

import (
	"database/sql"
	"os"


	_ "github.com/lib/pq"
)

func ConectDataBase() *sql.DB {

	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	sslmode := os.Getenv("DB_SSLMODE")

	conect := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", user, dbname, password, host, sslmode)
	db, err := sql.Open("postgres", conect)
	if err != nil {
		panic(err.Error())
	}
	return db
}
