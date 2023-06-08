package todoModel

import "github.com/henriquemdimer/go-crud/db"

func DeleteOne(id int64) (int64, error) {
	db, err := db.Open()
	if err != nil {
		return 0, nil
	}

	res, err := db.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		return 0, nil
	}

	return res.RowsAffected()
}
