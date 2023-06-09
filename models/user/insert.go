package userModel

import "github.com/henriquemdimer/go-crud/db"

func Insert(name string, password string) (id int64, err error) {
	db, err := db.Open()
	if err != nil {
		return
	}

	sql := "INSERT INTO users (name, password) VALUES ($1, $2) RETURNING id"
	err = db.QueryRow(sql, name, password).Scan(&id)

	db.Close()
	return
}
