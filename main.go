package main

import (
	"database/sql"
	"example/web-service-gin/src/controller"
	"example/web-service-gin/src/repository/postgres/repository"
	"example/web-service-gin/src/usecase"

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
	createAlbumUsecase := usecase.NewCreateAlbumUsecase(albumRepo)
	getAlbumUsecase := usecase.NewGetAlbumUsecase(albumRepo)
	listAlbumUsecase := usecase.NewListAlbumUsecase(albumRepo)
	updateAlbumUsecase := usecase.NewUpdateAlbumUsecase(albumRepo)
	deleteAlbumUsecase := usecase.NewDeleteAlbumUsecase(albumRepo)
	albumCon := controller.NewAlbumController(
		*createAlbumUsecase,
		*getAlbumUsecase,
		*listAlbumUsecase,
		*updateAlbumUsecase,
		*deleteAlbumUsecase,
	)

	router.GET("/albums", albumCon.ListAlbums)
	router.GET("/albums/:id", albumCon.GetAlbumByID)
	router.POST("/albums", albumCon.CreateAlbum)
	router.PUT("/albums", albumCon.UpdateAlbum)
	router.DELETE("/albums/:id", albumCon.DeleteAlbum)

	router.Run("localhost:8080")
}
