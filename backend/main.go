package main

import (
	"log"

	"github.com/sungyo4869/portfolio/db"
)

func main() {
	db, err := db.NewDB("root", "pass", "testdb", "mysql", "3306")
	if err != nil {
		log.Fatalln("main: err =", err)
	}

	defer db.Close()
}

