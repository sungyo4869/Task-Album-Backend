package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "embed"

	"github.com/go-sql-driver/mysql"
)

//go:embed schema.sql
var schema string

func NewDB(user, pass, dbName, host, port string) (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	c := mysql.Config{
		DBName:    dbName,
		User:      user,
		Passwd:    pass,
		Addr:      addr,
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, err
}
