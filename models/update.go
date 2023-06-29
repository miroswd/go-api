package models

import "github.com/miroswd/go-api/db"

func Update(id int64, todo Todo) (int64, error) {
	connection, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer connection.Close()

	res, err := connection.Exec(`UPDATE todos SET title=$1, description=$2, done=$3 WHERE id=$4`, todo.Title, todo.Description, todo.Done, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
