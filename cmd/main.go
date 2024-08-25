package main

import (
	"log"

	"github.com/bavithbhargav/go-album-backend/cmd/controllers"
	"github.com/bavithbhargav/go-album-backend/cmd/data"
	"github.com/gin-gonic/gin"
)

func main() {
	// Init inmemory albums
	data.InitAlbums()

	router := gin.Default()

	router.GET("/albums", controllers.GetAllAlbums)
	router.POST("/albums", controllers.CreateAlbum)
	router.GET("/albums/:id", controllers.GetAlbumById)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)
	router.PATCH("/albums", controllers.EditAlbum)

	log.Println("Server started at localhost:9090")
	router.Run("localhost:9090")
}
