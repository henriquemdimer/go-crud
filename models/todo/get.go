package todoModel

import "github.com/henriquemdimer/go-crud/db"

func GetOne(id int64) (todo Todo, err error) {
	db, err := db.Open()
	if err != nil {
		return
	}

	row := db.QueryRow("SELECT * FROM todos WHERE id=$1", id)
	err = row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)

	db.Close()
	return
}

func GetAll() (todos []Todo, err error) {
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
