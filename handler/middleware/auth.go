package middleware

import (
	"fmt"
	"net/http"
	"os"

	// "github.com/dgrijalva/jwt-go"
)

var secretKey = os.Getenv("SECRET_KEY")

func Auth(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			fmt.Println("認証情報ないよ")
			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
