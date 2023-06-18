package models

import "github.com/henriquemdimer/go-crud/db"

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func InsertUser(name string, password string) (id int64, err error) {
	db, err := db.Open()
	if err != nil {
		return
	}

	sql := "INSERT INTO users (name, password) VALUES ($1, $2) RETURNING id"
	err = db.QueryRow(sql, name, password).Scan(&id)

	db.Close()
	return
}

func GetUser(name string) (user User, err error) {
	db, err := db.Open()
	if err != nil {
		return
	}

	err = db.QueryRow("SELECT * FROM users WHERE name=$1", name).Scan(&user.Id, &user.Name, &user.Password)

	db.Close()
	return
}
