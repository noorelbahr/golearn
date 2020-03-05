package auth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/noorelbahr/golearn/helpers"
	"net/http"
	"time"
)

var signingKey = []byte("MyJWTSecretKey")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "ELBAHR"
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		panic(err.Error())
		return "", err
	}

	return tokenString, nil
}

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header["Authorization"] != nil {
			token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error.")
				}
				return signingKey, nil
			})

			if err != nil {
				helpers.JsonError(w, "Unauthorized. Error: " + err.Error(), 401)
				return
			}

			if token != nil && token.Valid {
				endpoint(w, r)
			}
		} else {
			helpers.JsonError(w, "Unauthorized.", 401)
		}
	})
}