package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "10.10.0.172"
	port     = 5432
	user     = "postgres"
	dbname   = "leetcode"
	password = "root"
)

func ConnectDB() (*sql.DB, error) {
	conn := fmt.Sprintf(`host=%s port=%d user=%s dbname=%s password=%s 
	sslmode=disable`, host, port, user, dbname, password)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return db, err
}
