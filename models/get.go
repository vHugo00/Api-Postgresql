package models

import (
	"github.com/vHugo00/Api-Postgresql/db"
)

func Get(id int64) (todo Todo, err error) {
	conn, err := db.Openconnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow("SELECT * FROM todos WHERE id = $1", id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Desctription, &todo.Done)
	return
}
