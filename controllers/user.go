package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"

	//"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/noorelbahr/golearn/helpers"
	"github.com/noorelbahr/golearn/models"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	users := models.AllUsers()

	helpers.JsonSuccess(w, users, 200)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	user, err := models.FindUser(id)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	helpers.JsonSuccess(w, user, 200)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	hash, _ := helpers.HashPassword(user.Password)
	user.Password = hash

	user, err = models.CreateUser(user)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	helpers.JsonSuccess(w, user, 201)
}

func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var user models.User
	user, err := models.FindUser(id)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	if user.Password != "" {
		hash, _ := helpers.HashPassword(user.Password)
		user.Password = hash
	}

	user, err = models.UpdateUser(user)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	helpers.JsonSuccess(w, user, 200)
}

func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	user, err := models.FindUser(id)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	err = models.DeleteUser(user)
	if err != nil {
		helpers.JsonError(w, err.Error(), 400)
		return
	}

	helpers.JsonSuccess(w, "User deleted successfully.", 204)
}
