package routes

import (
	"github.com/bavithbhargav/go-album-backend/cmd/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/albums", controllers.GetAllAlbums)
	router.POST("/albums", controllers.CreateAlbum)
	router.GET("/albums/:id", controllers.GetAlbumById)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)
	router.PATCH("/albums/:id", controllers.EditAlbum)
}
