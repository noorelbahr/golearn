package controllers

import (
	"encoding/json"
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

/**
 * Login
 */
func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest

	// Get request body
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	// Get user data by username
	user, err := models.FindUserByUsername(loginRequest.Username)
	if err != nil {
		helpers.JsonError(w, "Username is not registered.", 400)
		return
	}

	// Check user password
	isValid := helpers.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		helpers.JsonError(w, "Invalid password.", 400)
		return
	}

	// Generate JWT Token
	token, err := auth.GenerateJWT()
	if err != nil {
		helpers.JsonError(w, "Failed to generate Token.", 400)
		return
	}

	// Set response data
	var response LoginResponse
	response.Token = token

	helpers.JsonSuccess(w, response, 200)
}
