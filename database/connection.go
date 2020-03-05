package database

import (
	"github.com/jinzhu/gorm"
)

/**
 * Gorm Connect
 */
func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "golearn.db")
	//db, err := gorm.Open("mysql", "root:@/golearn?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return db
}