package main

import (
	"database/sql"
	"example/web-service-gin/src/domain"
	"example/web-service-gin/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func generateDB() (*sql.DB, error) {
	return sql.Open(
		"postgres",
		"host=postgresql dbname=go-demo user=go-demo password=password sslmode=disable")
}

func getAlbums(c *gin.Context) {
	db, dbErr := generateDB()
	if dbErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed database connection"})
		return
	}
	albumRepo := repository.NewAlbumRepository(db)
	album, err := albumRepo.FindAll(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func postAlbums(c *gin.Context) {
	db, dbErr := generateDB()
	if dbErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed database connection"})
		return
	}
	albumRepo := repository.NewAlbumRepository(db)
	var newAlbum domain.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed bind json"})
		return
	}
	err := albumRepo.Save(c, newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusCreated, "OK")
}

func getAlbumByID(c *gin.Context) {
	db, dbErr := generateDB()
	if dbErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed database connection"})
		return
	}
	albumRepo := repository.NewAlbumRepository(db)
	id := c.Param("id")
	album, err := albumRepo.FindById(c, id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed find data"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
