package controllers

import (
	"log"
	"net/http"

	"github.com/bavithbhargav/go-album-backend/cmd/database"
	"github.com/bavithbhargav/go-album-backend/cmd/models"
	"github.com/gin-gonic/gin"
)

func GetAllAlbums(c *gin.Context) {
	var albums []models.Album
	database.Instance.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func DeleteAlbum(c *gin.Context) {
	albumId := c.Param("id")

	if err := database.Instance.Delete(&models.Album{}, albumId).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Requested album not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Album deleted"})
}

func CreateAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Println("Error receiving http body", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error receiving http body"})
		return
	}

	database.Instance.Save(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumById(c *gin.Context) {
	albumId := c.Param("id")

	var album models.Album
	if err := database.Instance.First(&album, albumId).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Requested album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func EditAlbum(c *gin.Context) {
	id := c.Param("id")
	var album models.Album

	if err := database.Instance.First(&album, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	if err := c.ShouldBindJSON(&album); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Instance.Save(&album)
	c.IndentedJSON(http.StatusOK, album)
}
