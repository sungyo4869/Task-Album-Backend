package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sungyo4869/portfolio/model"
	"github.com/sungyo4869/portfolio/service"
)

type LoginHandler struct{
	svc *service.UserService
}

func NewLoginHandler(svc *service.UserService) *LoginHandler {
	return &LoginHandler{
		svc: svc,
	}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var req model.LoginRequest
		var res model.LoginResponse

		params := r.URL.Query()
		// クエリパラメータの値を取得
		pramStr := params.Get("email")
		if pramStr != "" {
			req.Email = pramStr
		} else {
			http.Error(w, "Authorization email is missing", http.StatusUnauthorized)
			log.Println("Login: Failed to get parameters")
			return
		}

		pramStr = params.Get("password")
		if pramStr != "" {
			req.Pass = pramStr
		} else {
			http.Error(w, "Authorization password is missing", http.StatusUnauthorized)
			log.Println("Login: Failed to get parameters")
			return
		}

		// DBに確認
		user, err := h.svc.ReadUser(r.Context(), req.Email, req.Pass)
		if err != nil {
			http.Error(w, "Email or password is missing", http.StatusUnauthorized)
			log.Println("Login: Authentication failed, err = ", err)
			return
		} 

		token, err := h.GenerateToken(user.ID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Login: Filed to generate token, err = ", err)
			return
		}

		res.Token = token

		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Login: Filed to encoding json, err = ", err)
			return
		}
	}
}

var secretKey = os.Getenv("SECRET_KEY")

func (h *LoginHandler) GenerateToken(uid int64) (string, error) {

	expirationTime := time.Now().Add(time.Hour * 1).Unix()

	claims := jwt.MapClaims{
		"user_id": uid,
		"exp":     expirationTime, // トークンの有効期限（1時間）
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// secretKey を使ってトークンを署名
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
