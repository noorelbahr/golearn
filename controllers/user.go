package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/noorelbahr/golearn/helpers"
	"github.com/noorelbahr/golearn/models"
	"net/http"
	"strconv"
)

/**
 * Get All Users
 */
func AllUsers(w http.ResponseWriter, r *http.Request) {
	// Get all users
	users := models.AllUsers()

	helpers.JsonSuccess(w, users, 200)
}

/**
 * Get User Detail
 */
func FindUser(w http.ResponseWriter, r *http.Request) {
	// Get param id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Get user data
	user, err := models.FindUser(id)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	helpers.JsonSuccess(w, user, 200)
}

/**
 * Create User
 */
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Get request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	// Hash user password
	hash, _ := helpers.HashPassword(user.Password)
	user.Password = hash

	// Create user data
	user, err = models.CreateUser(user)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	helpers.JsonSuccess(w, user, 201)
}

/**
 * Update User
 */
func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	// Get param id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Get user data
	user, err := models.FindUser(id)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	// Get request body
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	// Update password if needed
	if user.Password != "" {
		hash, _ := helpers.HashPassword(user.Password)
		user.Password = hash
	}

	// Update user data
	user, err = models.UpdateUser(user)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	helpers.JsonSuccess(w, user, 200)
}

/**
 * Delete User
 */
func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	// Get param id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Get user data
	user, err := models.FindUser(id)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	// Delete user data
	err = models.DeleteUser(user)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	helpers.JsonSuccess(w, "User deleted successfully.", 204)
}
