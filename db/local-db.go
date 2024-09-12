package db

import (
	"database/sql"
	"log"
)

func NewDB() (*sql.DB, error) {
	dsn := "root:2PL|RQ)hpevE@tcp(127.0.0.1:3306)/testdb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}
	// defer db.Close()

	return db, nil
}
