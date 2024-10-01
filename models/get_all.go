package models

import (
	"github.com/vHugo00/Api-Postgresql/db"
)

func GetAll() (todos []Todo, err error) {
	conn, err := db.Openconnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM todos ORDER BY id")
	if err != nil {
		return
	}
	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Desctription, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}
	return
}
