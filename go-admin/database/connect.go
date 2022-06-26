package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {

	_, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		// panic is thowring an error
		panic("failed to connect database")
	}

}
