package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/sungyo4869/portfolio/db"
	"github.com/sungyo4869/portfolio/handler/router"
)

func main() {
	realMain()
}

func realMain() {

	err := godotenv.Load()
	if err != nil {
		log.Print(err)
		return
	}

	port := os.Getenv("PORT")

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPath := os.Getenv("DB_PATH")
	dbPort := os.Getenv("DB_PORT")

	db, err := db.NewDB(dbUser, dbPass, dbName, dbPath, dbPort)
	if err != nil {
		log.Fatalln("main: err =", err)
	}
	defer db.Close()

	mux := router.NewRouter(db)

	srv := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	srv.ListenAndServe()

}
