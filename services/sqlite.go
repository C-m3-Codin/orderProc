package services

import (
	"c-m3-codin/ordProc/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnections(db string) (dbobj *gorm.DB) {
	if db == "sqlite3" {
		dbobj = getGormObj()
	}

	runMigrations(dbobj)
	return
}

func getGormObj() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return
}

func runMigrations(dbobj *gorm.DB) {
	dbobj.AutoMigrate(models.Order{})
	return
}
