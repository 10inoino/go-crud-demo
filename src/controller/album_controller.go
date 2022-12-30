package controller

import (
	"example/web-service-gin/src/domain"
	repository_interface "example/web-service-gin/src/repository/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	albumRepo repository_interface.AlbumRepository
}

func NewAlbumController(
	albumRepo repository_interface.AlbumRepository,
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
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, album)
}

func (con *AlbumController) CreateAlbum(ctx *gin.Context) {
	// TODO:リクエストの型で受ける
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

func (con *AlbumController) DeleteAlbum(ctx *gin.Context) {
	id := ctx.Param("id")
	err := con.albumRepo.DeleteById(ctx, id)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, "OK")
}

func (con *AlbumController) UpdateAlbum(ctx *gin.Context) {
	// TODO:リクエストの型で受ける
	var newAlbum domain.Album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed bind json"})
		return
	}
	err := con.albumRepo.Update(ctx, newAlbum)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, "OK")
}
