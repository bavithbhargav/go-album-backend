package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bavithbhargav/go-album-backend/cmd/database"
	routers "github.com/bavithbhargav/go-album-backend/cmd/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load envs
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		panic("Error loading envs")
	}

	// Connect to DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	database.Connect(dsn)

	// Migrate DB
	database.Migrate()

	r := gin.Default()
	routers.RegisterRoutes(r)

	log.Println("Server started at localhost:9090")
	r.Run("localhost:9090")
}
