package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Open() (db *sql.DB, err error) {
	godotenv.Load()

	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	dbname := os.Getenv("DBNAME")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")

	str := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)
	db, err = sql.Open("postgres", str)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	return
}
