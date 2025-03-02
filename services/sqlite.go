package services

import (
	"c-m3-codin/ordProc/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect to the database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting *sql.DB object: %v", err)
	}

	// Set connection pool parameters
	sqlDB.SetMaxOpenConns(100)                 // Maximum number of open connections
	sqlDB.SetMaxIdleConns(20)                  // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(10 * time.Minute) // Maximum time a connection can be reused

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
