package db

import (
	"log"

	"github.com/jeroensmink98/gobook/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	return db
}

func InitDb(db *gorm.DB) {
	err := db.AutoMigrate(&models.Book{}, &models.Genre{})
	if err != nil {
		log.Fatalf("error migrating table: %v", err)
	}
}

func Delete(db *gorm.DB) {
	err := db.Migrator().DropTable(&models.Book{}, &models.Genre{})
	if err != nil {
		log.Fatalf("Could not delete records: %v", err)
	}
}

func SeedDb(db *gorm.DB) {
	// seed database
}
