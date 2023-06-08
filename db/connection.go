package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Open() (db *sql.DB, err error) {
	str := "host=localhost port=5432 user=root password=admin dbname=main sslmode=disable"
	db, err = sql.Open("postgres", str)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	return
}
