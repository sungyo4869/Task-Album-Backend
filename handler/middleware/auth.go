package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = os.Getenv("SECRET_KEY")

type UserIDKey struct{}

func Auth(h http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")

        if tokenString == "" {
            http.Error(w, "認証情報ないよ", http.StatusUnauthorized)
            return
        }
		log.Println("token = ", tokenString)

        // Remove the "Bearer " prefix if it exists
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		log.Println("token = ", tokenString)

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(secretKey), nil
        })

        if err != nil {
            http.Error(w, "トークンが無効です", http.StatusUnauthorized)
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Println("claims[user_id] = ", claims["user_id"])

            userID, ok := claims["user_id"].(float64)
            if !ok {
                http.Error(w, "ユーザーIDが見つかりません", http.StatusUnauthorized)
				log.Println("middleware: ユーザーIDが見つかりません")
                return
            }

            // Store the user ID in the request context
            ctx := context.WithValue(r.Context(), UserIDKey{}, int64(userID))
            r = r.WithContext(ctx)

            h.ServeHTTP(w, r)
        } else {
            http.Error(w, "トークンが無効ですwi", http.StatusUnauthorized)
            return
        }
    }

    return http.HandlerFunc(fn)
}
