package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/noorelbahr/golearn/helpers"
	"github.com/noorelbahr/golearn/models"
	"io/ioutil"
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

	// Get request body -> multipart/form-data
	err := r.ParseMultipartForm(2 << 20)
	if err != nil {
		helpers.JsonError(w, "Err: " + err.Error(), 400)
		return
	}

	// Hash user password
	hash, _ := helpers.HashPassword(user.Password)

	// Set user data
	user.Username = r.PostForm.Get("username")
	user.Fullname = r.PostForm.Get("fullname")
	user.Password = hash

	// Check user file
	file, handler, err := r.FormFile("picture")
	if err == nil {
		defer file.Close()

		// Make temp file
		tempFile, err := ioutil.TempFile("assets", "*-" + handler.Filename)
		if err != nil {
			helpers.JsonError(w, "Err3: " + err.Error(), 400)
			return
		}
		defer tempFile.Close()

		// Read the file
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			helpers.JsonError(w, "Err4: " + err.Error(), 400)
			return
		}

		// Save file
		_, err = tempFile.Write(fileBytes)
		if err == nil {
			user.Picture = tempFile.Name()
		}
	}

	// Create user data
	user, err = models.CreateUser(user)
	if err != nil {
		helpers.JsonError(w, "Err5: " + err.Error(), 400)
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
