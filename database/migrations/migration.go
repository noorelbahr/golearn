package migrations

import (
	"fmt"
	"github.com/noorelbahr/golearn/database"
	"github.com/noorelbahr/golearn/helpers"
	"github.com/noorelbahr/golearn/models"
)

/**
 * Set Initial Migration and Seeder
 */
func InitialMigration() {
	db := database.Connect()
	defer db.Close()

	// Migrate
	err := db.AutoMigrate(&models.User{}).Error
	if err == nil {
		fmt.Println("User migration: OK")
	}

	// Create default user data
	_, err = models.FindUserByUsername("johndoe")
	if err != nil {
		hash, _ := helpers.HashPassword("123123")

		var user models.User
		user.Username = "johndoe"
		user.Password = hash
		user.Fullname = "John Doe"
		_, err = models.CreateUser(user)
		if err == nil {
			fmt.Println("User seeder: OK")
		}
	}
}
