package database

import (
	"log"

	"github.com/bavithbhargav/go-album-backend/cmd/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to DB")
}

func Migrate() {
	Instance.AutoMigrate(&models.Album{})
	log.Println("Database migration completed")
}
