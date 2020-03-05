package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/noorelbahr/golearn/helpers"
	"time"
)

var db *gorm.DB
var err error

type User struct {
	ID        	uint 		`gorm:"primary_key" json:"id"`
	Username 	string 		`gorm:"column:username;unique_index" json:"username"`
	Password 	string 		`gorm:"column:password;size:255" json:"password"`
	Fullname 	string 		`gorm:"column:fullname" json:"fullname"`
	Picture 	string 		`gorm:"column:picture" json:"picture"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
	DeletedAt 	*time.Time 	`sql:"index" json:"deleted_at"`
}

type Users []User

func InitialMigration() {
	db := connect()
	defer db.Close()

	// Migrate
	err := db.AutoMigrate(&User{}).Error
	if err == nil {
		fmt.Println("User migration: OK")
	}

	// Create default user data
	_, err = FindUserByUsername("johndoe")
	if err != nil {
		hash, _ := helpers.HashPassword("123123")

		var user User
		user.Username = "johndoe"
		user.Password = hash
		user.Fullname = "John Doe"
		_, err = CreateUser(user)
		if err == nil {
			fmt.Println("User seeder: OK")
		}
	}
}

func AllUsers() []User {
	db := connect()
	defer db.Close()

	var users []User
	db.Find(&users)

	return users
}

func FindUser(id int) (User, error) {
	db := connect()
	defer db.Close()

	var user User
	err := db.First(&user, id).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func FindUserByUsername(username string) (User, error) {
	db := connect()
	defer db.Close()

	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func CreateUser(user User) (User, error) {
	db := connect()
	defer db.Close()

	err := db.Create(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func UpdateUser(user User) (User, error) {
	db := connect()
	defer db.Close()

	err := db.Model(&user).Updates(user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func DeleteUser(user User) error {
	db := connect()
	defer db.Close()

	return db.Delete(&user).Error
}

/**
 * Gorm Connect
 */
func connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "golearn.db")
	//db, err := gorm.Open("mysql", "root:@/golearn?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return db
}
