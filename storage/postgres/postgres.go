package postgres

import (
	"database/sql"
	_"github.com/lib/pq"
)


func ConnectDB() (*sql.DB, error){
	url := "postgresql://postgres:root@localhost/leetcode"
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}