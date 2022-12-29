package controller

import (
	"example/web-service-gin/src/domain"
	"example/web-service-gin/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	albumRepo repository.AlbumRepository
}

func NewAlbumController(
	albumRepo repository.AlbumRepository,
) *AlbumController {
	return &AlbumController{
		albumRepo: albumRepo,
	}
}

func (con *AlbumController) GetAlbums(ctx *gin.Context) {
	album, err := con.albumRepo.FindAll(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, album)
}

func (con *AlbumController) GetAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")
	album, err := con.albumRepo.FindById(ctx, id)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed find data"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, album)
}

func (con *AlbumController) PostAlbum(ctx *gin.Context) {
	var newAlbum domain.Album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed bind json"})
		return
	}
	err := con.albumRepo.Save(ctx, newAlbum)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, "OK")
}
