package main

import (
	"database/sql"
	"example/web-service-gin/db/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  int    `json:"price"`
}

func getAlbums(c *gin.Context) {
	db, dbErr := sql.Open(
		"postgres",
		"host=postgresql dbname=go-demo user=go-demo password=password sslmode=disable")
	if dbErr != nil {
		fmt.Print(dbErr)
		return
	}
	m, err := models.Albums().All(c, db)
	if err != nil {
		fmt.Print(err)
		return
	}

	c.IndentedJSON(http.StatusOK, m)
}

func postAlbums(c *gin.Context) {
	db, dbErr := sql.Open(
		"postgres",
		"host=postgresql dbname=go-demo user=go-demo password=password sslmode=disable")
	if dbErr != nil {
		fmt.Print(dbErr)
	}
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	album := &models.Album{
		ID:     newAlbum.ID,
		Title:  newAlbum.Title,
		Artist: newAlbum.Artist,
		Price:  newAlbum.Price,
	}

	err := album.Insert(c, db, boil.Infer())
	if err != nil {
		fmt.Print(err)
	}

	c.IndentedJSON(http.StatusCreated, album)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	db, dbErr := sql.Open(
		"postgres",
		"host=postgresql dbname=go-demo user=go-demo password=password sslmode=disable")
	if dbErr != nil {
		fmt.Print(dbErr)
	}
	m, err := models.Albums(
		qm.Where("id=?", id),
	).One(c, db)
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, m)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
