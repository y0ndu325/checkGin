package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Storage interface {
	Create() album
	Read() album
	ReadOne() (album, error)
	Update() album
	Delete() album
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type MemoryStorage struct {
	albums []album
}

var storage = NewMemoryStorage()

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func (s MemoryStorage) Create(a album) album {
	s.albums = append(s.albums, a)
	return a
}

func (s MemoryStorage) ReadOne(id string) (album, error) {

	for _, a := range s.albums {
		if a.ID == id {
			return a, nil
		}
	}
	return album{}, errors.New("not found")
}

func (s MemoryStorage) Read() []album {
	return s.albums
}

func (s MemoryStorage) Update(id string, a album) (album, error) {

	for i := range albums {
		if s.albums[i].ID == id {
			return s.albums[i], nil
		}
	}
	return album{}, errors.New("not found")
}

func (s MemoryStorage) Delete(id string) error {
	for i, a := range s.albums {
		if a.ID == id {
			s.albums = append(s.albums[:i], s.albums[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func NewMemoryStorage() MemoryStorage {
	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	return MemoryStorage{albums: albums}
}

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
