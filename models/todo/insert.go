package todoModel

import "github.com/henriquemdimer/go-crud/db"

func InsertOne(todo Todo) (id int64, err error) {
	db, err := db.Open()
	if err != nil {
		return
	}

	sql := "INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id"
	err = db.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	db.Close()
	return
}
