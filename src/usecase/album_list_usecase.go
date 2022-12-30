package usecase

import (
	"example/web-service-gin/src/domain"
	repository_interface "example/web-service-gin/src/repository/interface"

	"github.com/gin-gonic/gin"
)

type AlbumListUsecase struct {
	albumRepo repository_interface.AlbumRepository
}

func NewAlbumListUsecase(
	albumRepo repository_interface.AlbumRepository,
) *AlbumListUsecase {
	return &AlbumListUsecase{
		albumRepo: albumRepo,
	}
}

func (usecase *AlbumListUsecase) Execute(
	ctx *gin.Context,
) (*[]domain.Album, error) {
	albums, err := usecase.albumRepo.FindAll(ctx)
	if err != nil {
		return nil, domain.NewNotFoundError("Failed list albums")
	}
	return albums, nil
}
