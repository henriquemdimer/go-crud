package todoModel

import "github.com/henriquemdimer/go-crud/db"

func UpdateOne(id int64, todo Todo) (int64, error) {
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
