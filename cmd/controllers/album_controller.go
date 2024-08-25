package controllers

import (
	"log"
	"net/http"

	"github.com/bavithbhargav/go-album-backend/cmd/data"
	"github.com/bavithbhargav/go-album-backend/cmd/models"
	"github.com/bavithbhargav/go-album-backend/cmd/utils"
	"github.com/gin-gonic/gin"
)

func GetAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Albums)
}

func DeleteAlbum(c *gin.Context) {
	albumId := c.Param("id")

	for i, album := range data.Albums {
		if album.ID == albumId {
			data.Albums = append(data.Albums[:i], data.Albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, data.Albums)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Requested album not found"})
}

func CreateAlbum(c *gin.Context) {
	var newAlbum models.Album
	newAlbum.ID = utils.Random3DigitString()

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Println("Error receiving http body", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error receiving http body"})
		return
	}

	data.Albums = append(data.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumById(c *gin.Context) {
	albumId := c.Param("id")

	for i, album := range data.Albums {
		if album.ID == albumId {
			c.IndentedJSON(http.StatusOK, data.Albums[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Requested album not found"})
}

func EditAlbum(c *gin.Context) {
	var updateAlbum models.Album

	if err := c.BindJSON(&updateAlbum); err != nil {
		log.Println("Error receiving http body", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error receiving http body"})
		return
	}

	for i, album := range data.Albums {
		if album.ID == updateAlbum.ID {
			data.Albums[i] = updateAlbum
			c.IndentedJSON(http.StatusOK, data.Albums[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Requested album not found"})
}
