package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Open() error {
	var err error
	Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	db, err := Db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}
