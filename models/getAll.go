package models

import "github.com/miroswd/go-api/db"

func GetAll() (todos []Todo, err error) {
	connection, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer connection.Close()

	rows, err := connection.Query(`SELECT * FROM todos`)

	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
