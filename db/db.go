package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

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

	return db, err
}
