package models

import "github.com/miroswd/go-api/db"

func Insert(todo Todo) (id int64, err error) {

	connection, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer connection.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = connection.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
