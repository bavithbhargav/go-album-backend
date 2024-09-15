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

	_, exists := checkIfAlbumExists(albumId)
	if !exists {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Requested album not found"})
		return
	}

	database.Instance.Delete(&models.Album{}, albumId)
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

	album, exists := checkIfAlbumExists(albumId)
	if !exists {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Requested album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, *album)
}

func EditAlbum(c *gin.Context) {
	var updateAlbum models.Album

	if err := c.ShouldBindJSON(&updateAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, exists := checkIfAlbumExists(updateAlbum.ID)
	if !exists {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	database.Instance.Save(&updateAlbum)
	c.IndentedJSON(http.StatusOK, updateAlbum)
}

func checkIfAlbumExists(albumId any) (*models.Album, bool) {
	var album models.Album

	database.Instance.First(&album, albumId)
	if album.ID == 0 {
		return &models.Album{}, false
	}

	return &album, true
}
