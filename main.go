package main

import (
	"database/sql"
	"example/web-service-gin/src/controller"
	"example/web-service-gin/src/repository/postgres/repository"

	"github.com/gin-gonic/gin"
)

func generateDB() (*sql.DB, error) {
	return sql.Open(
		"postgres",
		"host=postgresql dbname=go-demo user=go-demo password=password sslmode=disable")
}

func main() {
	router := gin.Default()
	db, dbErr := generateDB()
	if dbErr != nil {
		panic("failed database connection")
	}
	albumRepo := repository.NewAlbumRepository(db)
	albumCon := controller.NewAlbumController(albumRepo)

	router.GET("/albums", albumCon.GetAlbums)
	router.GET("/albums/:id", albumCon.GetAlbumByID)
	router.POST("/albums", albumCon.CreateAlbum)
	router.PUT("/albums", albumCon.UpdateAlbum)
	router.DELETE("/albums/:id", albumCon.DeleteAlbum)

	router.Run("localhost:8080")
}
