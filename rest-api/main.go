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

//getAlbums returns all albums as JSON

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) //serialized JSON and status code
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

var albums = []album{
	{ID: "1", Title: "Sweet O child of mine", Artist: "Guns and roses", Price: 60.00},
	{ID: "2", Title: "One more Light", Artist: "Linkin Park", Price: 20.00},
	{ID: "3", Title: "Californication", Artist: "Red Hot Chilli Peppers", Price: 30.00},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
