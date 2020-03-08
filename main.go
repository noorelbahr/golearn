package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/noorelbahr/golearn/auth"
	"github.com/noorelbahr/golearn/controllers"
	"github.com/noorelbahr/golearn/database/migrations"
	"github.com/noorelbahr/golearn/helpers"
	"log"
	"net/http"
)

/**
 * Handle Requests or Routes
 */
func handleRequests() {
	// Register routes
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.Handle("/", auth.IsAuthorized(controllers.HomePage)).Methods("GET")
	router.Handle("/users", auth.IsAuthorized(controllers.AllUsers)).Methods("GET")
	router.Handle("/users/{id}", auth.IsAuthorized(controllers.FindUser)).Methods("GET")
	router.Handle("/users", auth.IsAuthorized(controllers.CreateUser)).Methods("POST")
	router.Handle("/users/{id}", auth.IsAuthorized(controllers.UpdateUser)).Methods("PUT")
	router.Handle("/users/{id}", auth.IsAuthorized(controllers.DeleteUser)).Methods("DELETE")

	// Handle static file
	fs := http.FileServer(helpers.MyFS{Dir: http.Dir("./assets/")})
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs)).Methods("GET")

	// Lesten and serve on 8082
	log.Fatal(http.ListenAndServe(":8082", router))
}

/**
 * Main Function
 */
func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Run migration
	migrations.InitialMigration()

	// Handle requests
	handleRequests()
}