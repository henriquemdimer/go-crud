package models

import "github.com/henriquemdimer/go-crud/db"

type Todo struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func GetOneTodo(id int64) (todo Todo, err error) {
	db, err := db.Open()
	if err != nil {
		return
	}

	row := db.QueryRow("SELECT * FROM todos WHERE id=$1", id)
	err = row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)

	db.Close()
	return
}

func GetAllTodos() (todos []Todo, err error) {
	db, err := db.Open()
	if err != nil {
		return
	}

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	db.Close()
	return
}

func InsertOneTodo(todo Todo) (id int64, err error) {
	db, err := db.Open()
	if err != nil {
		return
	}

	sql := "INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id"
	err = db.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	db.Close()
	return
}

func UpdateOneTodo(id int64, todo Todo) (int64, error) {
	db, err := db.Open()
	if err != nil {
		return 0, err
	}

	res, err := db.Exec(`UPDATE todos SET title=$1, description=$2, done=$3 WHERE id=$4`, todo.Title, todo.Description, todo.Done, id)
	if err != nil {
		return 0, err
	}

	db.Close()
	return res.RowsAffected()
}

func DeleteOneTodo(id int64) (int64, error) {
	db, err := db.Open()
	if err != nil {
		return 0, err
	}

	res, err := db.Exec(`DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	db.Close()
	return res.RowsAffected()
}
