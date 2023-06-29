package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/miroswd/go-api/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	stringConnection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disabled",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database,
	)

	connection, err := sql.Open("postgres", stringConnection)

	if err != nil {
		panic(err)
	}

	err = connection.Ping() // check if the connection is really connected -> select 1

	return connection, err

}
