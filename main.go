package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var storage = NewStorage()

func main() {
	router := getRouter()
	router.Run(":8080")
}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.PUT("/albums/:id", updateAlbumByID)
	return router
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, storage.Read())
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, err := storage.ReadOne(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	storage.Create(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	err := storage.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "not found"})
		return
	}
	c.IndentedJSON(http.StatusNoContent, album{})
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var newAlbum album

	c.BindJSON(&newAlbum)
	album, err := storage.Update(id, newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}
