package services

import (
	"c-m3-codin/ordProc/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnections(db string) (dbobj *gorm.DB) {
	if db == "sqlite3" {
		dbobj = getGormObjSQLITE()
	} else if db == "postgres" {
		dbobj = getGormObjpostgres()
	}

	runMigrations(dbobj)
	return
}

func getGormObjpostgres() (db *gorm.DB) {
	// Define PostgreSQL connection details
	dsn := "host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable"

	// Open a connection to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("failed to connect to the database")
	}
	return
}

func getGormObjSQLITE() (db *gorm.DB) {
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
