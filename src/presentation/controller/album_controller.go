package controller

import (
	"example/web-service-gin/src/domain/request"
	"example/web-service-gin/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	createAlbumUsecase usecase.CreateAlbumUsecase
	getAlbumUsecase    usecase.GetAlbumUsecase
	listAlbumUsecase   usecase.ListAlbumUsecase
	updateAlbumUsecase usecase.UpdateAlbumUsecase
	deleteAlbumUsecase usecase.DeleteAlbumUsecase
}

func NewAlbumController(
	createAlbumUsecase usecase.CreateAlbumUsecase,
	getAlbumUsecase usecase.GetAlbumUsecase,
	listAlbumUsecase usecase.ListAlbumUsecase,
	updateAlbumUsecase usecase.UpdateAlbumUsecase,
	deleteAlbumUsecase usecase.DeleteAlbumUsecase,
) *AlbumController {
	return &AlbumController{
		createAlbumUsecase: createAlbumUsecase,
		getAlbumUsecase:    getAlbumUsecase,
		listAlbumUsecase:   listAlbumUsecase,
		updateAlbumUsecase: updateAlbumUsecase,
		deleteAlbumUsecase: deleteAlbumUsecase,
	}
}

func (con *AlbumController) ListAlbums(ctx *gin.Context) {
	albums, err := con.listAlbumUsecase.Execute(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, albums)
}

func (con *AlbumController) GetAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")
	album, err := con.getAlbumUsecase.Execute(ctx, id)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, album)
}

func (con *AlbumController) CreateAlbum(ctx *gin.Context) {
	var newAlbum request.AlbumCreateRequest
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed bind json"})
		return
	}
	err := con.createAlbumUsecase.Execute(
		ctx,
		newAlbum.ID,
		newAlbum.Title,
		newAlbum.Artist,
		newAlbum.Price,
	)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, "OK")
}

func (con *AlbumController) DeleteAlbum(ctx *gin.Context) {
	id := ctx.Param("id")
	err := con.deleteAlbumUsecase.Execute(ctx, id)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, "OK")
}

func (con *AlbumController) UpdateAlbum(ctx *gin.Context) {
	var newAlbum request.AlbumUpdateRequest
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed bind json"})
		return
	}
	err := con.updateAlbumUsecase.Execute(
		ctx,
		newAlbum.ID,
		newAlbum.Title,
		newAlbum.Artist,
		newAlbum.Price,
	)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, "OK")
}
