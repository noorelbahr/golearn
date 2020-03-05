package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/noorelbahr/golearn/auth"
	"github.com/noorelbahr/golearn/controllers"
	"github.com/noorelbahr/golearn/models"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("This is my homepage")
}



func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.Handle("/", auth.IsAuthorized(homePage)).Methods("GET")
	router.Handle("/users", auth.IsAuthorized(controllers.AllUsers)).Methods("GET")
	router.Handle("/users/{id}", auth.IsAuthorized(controllers.FindUser)).Methods("GET")
	router.Handle("/users", auth.IsAuthorized(controllers.CreateUser)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main()  {
	models.InitialMigration()

	handleRequests()
}