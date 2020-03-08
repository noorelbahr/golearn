package database

import (
	"github.com/jinzhu/gorm"
	"os"
)

/**
 * DB Connect
 */
func Connect() *gorm.DB {
	dialect := os.Getenv("DB_DIALECT")
	args := os.Getenv("DB_NAME")
	// args := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(dialect, args)
	if err != nil {
		panic(err.Error())
	}
	return db
}