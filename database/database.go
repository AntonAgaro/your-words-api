package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func ConnectDb(dsn string) {
	if len(dsn) == 0 {
		log.Println("DATABASE_URL is not set")
		return
	}
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot connect to db with dsn %v", dsn)
	}
	log.Printf("Connected to DB with dsn %v", dsn)
}
