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

	db, err := db.NewDB()
	if err != nil {
		log.Fatalln("main: err =", err)
	} else {
		log.Println("errなかったぽ")
		err = db.Ping()
		if err != nil {
			log.Println("pingできてないよ, err = ", err)
		} else {
			log.Println("pingできたっぽ")
		}
	}
	defer db.Close()

	mux := router.NewRouter(db)

	srv := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	srv.ListenAndServe()

}
