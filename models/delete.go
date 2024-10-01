package models

import "github.com/vHugo00/Api-Postgresql/db"

func Delete(id int64) (int64, error) {
	conn, err := db.Openconnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
