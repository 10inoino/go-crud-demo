package repository_interface

import (
	"example/web-service-gin/src/domain"

	"github.com/gin-gonic/gin"
)

type AlbumRepository interface {
	Save(ctx *gin.Context, album domain.Album) error
	FindAll(ctx *gin.Context) (*[]domain.Album, error)
	FindById(ctx *gin.Context, id string) (*domain.Album, error)
	DeleteById(ctx *gin.Context, id string) error
	Update(ctx *gin.Context, album domain.Album) error
}
