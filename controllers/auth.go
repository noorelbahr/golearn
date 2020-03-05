package controllers

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/noorelbahr/golearn/auth"
	"github.com/noorelbahr/golearn/helpers"
	"github.com/noorelbahr/golearn/models"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	user := models.FindUserByUsername(loginRequest.Username)

	isValid := helpers.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		helpers.JsonError(w, "Invalid password.", 400)
		return
	}

	token, err := auth.GenerateJWT()
	if err != nil {
		helpers.JsonError(w, "Failed to generate Token.", 400)
		return
	}

	var response LoginResponse
	response.Token = token

	helpers.JsonSuccess(w, response, 200)
}

/**
 * Gorm Connect
 */
func connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "golearn.db")
	if err != nil {
		panic(err.Error())
	}
	return db
}
