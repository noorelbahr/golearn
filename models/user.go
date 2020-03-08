package models

import (
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/noorelbahr/golearn/database"
	"time"
)

type User struct {
	ID        	uint 		`gorm:"primary_key" json:"id"`
	Username 	string 		`gorm:"column:username;unique_index" json:"username"`
	Password 	string 		`gorm:"column:password;size:255" json:"password"`
	Fullname 	string 		`gorm:"column:fullname" json:"fullname"`
	Picture 	string 		`gorm:"column:picture" json:"picture"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
	DeletedAt 	*time.Time 	`sql:"index" json:"-"`
}

/**
 * Custom User's MarshalJSON function
 */
func (u User) MarshalJSON() ([]byte, error) {
	type Alias User

	// Set user picture url
	pictureUrl := u.Picture
	if pictureUrl != "" {
		pictureUrl = "http://localhost:8082/" + pictureUrl
	}

	return json.Marshal(&struct {
		Alias
		Picture string `json:"picture"`
	}{
		Alias: (Alias)(u),
		Picture: pictureUrl,
	})
}

func AllUsers() []User {
	db := database.Connect()
	defer db.Close()

	var users []User
	db.Find(&users)

	return users
}

func FindUser(id int) (User, error) {
	db := database.Connect()
	defer db.Close()

	var user User
	err := db.First(&user, id).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func FindUserByUsername(username string) (User, error) {
	db := database.Connect()
	defer db.Close()

	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func CreateUser(user User) (User, error) {
	db := database.Connect()
	defer db.Close()

	err := db.Create(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func UpdateUser(user User) (User, error) {
	db := database.Connect()
	defer db.Close()

	err := db.Model(&user).Updates(user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func DeleteUser(user User) error {
	db := database.Connect()
	defer db.Close()

	return db.Delete(&user).Error
}
