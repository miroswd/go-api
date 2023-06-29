package models

import "github.com/miroswd/go-api/db"

func Update(id int64) (int64, error) {
	connection, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer connection.Close()

	res, err := connection.Exec(`DELETE FROM todos WHERE id=$1`, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
