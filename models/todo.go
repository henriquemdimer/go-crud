package models

import "github.com/henriquemdimer/go-crud/db"

type Todo struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	UserID      int64  `json:"user_id"`
}

func GetOneTodo(id int64, userID int64) (todo Todo, err error) {
	db, err := db.Open()
	if err != nil {
		return
	}

	row := db.QueryRow("SELECT * FROM todos WHERE id=$1 AND user_id=$2", id, userID)
	err = row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done, &todo.UserID)

	db.Close()
	return
}

func GetAllTodos(userID int64) (todos []Todo, err error) {
	db, err := db.Open()
	if err != nil {
		return
	}

	rows, err := db.Query("SELECT * FROM todos WHERE user_id=$1", userID)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done, &todo.UserID)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	db.Close()
	return
}

func InsertOneTodo(todo Todo, userID int64) (id int64, err error) {
	db, err := db.Open()
	if err != nil {
		return
	}

	sql := "INSERT INTO todos (title, description, done, user_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err = db.QueryRow(sql, todo.Title, todo.Description, todo.Done, userID).Scan(&id)

	db.Close()
	return
}

func UpdateOneTodo(id int64, todo Todo, userID int64) (int64, error) {
	db, err := db.Open()
	if err != nil {
		return 0, err
	}

	res, err := db.Exec(`UPDATE todos SET title=$1, description=$2, done=$3 WHERE id=$4 AND user_id=$5`, todo.Title, todo.Description, todo.Done, id, userID)
	if err != nil {
		return 0, err
	}

	db.Close()
	return res.RowsAffected()
}

func DeleteOneTodo(id int64, userID int64) (int64, error) {
	db, err := db.Open()
	if err != nil {
		return 0, err
	}

	res, err := db.Exec(`DELETE FROM todos WHERE id=$1 AND user_id=$2`, id, userID)
	if err != nil {
		return 0, err
	}

	db.Close()
	return res.RowsAffected()
}
