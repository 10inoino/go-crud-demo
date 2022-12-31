package usecase

import (
	"example/web-service-gin/src/domain"
	repository_interface "example/web-service-gin/src/repository/interface"

	"github.com/gin-gonic/gin"
)

type ListAlbumUsecase struct {
	albumRepo repository_interface.AlbumRepository
}

func NewListAlbumUsecase(
	albumRepo repository_interface.AlbumRepository,
) *ListAlbumUsecase {
	return &ListAlbumUsecase{
		albumRepo: albumRepo,
	}
}

func (usecase *ListAlbumUsecase) Execute(
	ctx *gin.Context,
) (*[]domain.Album, error) {
	albums, err := usecase.albumRepo.FindAll(ctx)
	if err != nil {
		return nil, domain.NewNotFoundError("Failed list albums")
	}
	return albums, nil
}
