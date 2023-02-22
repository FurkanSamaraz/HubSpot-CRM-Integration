package configs

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetCollection() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
