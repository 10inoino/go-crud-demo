package usecase

import (
	"example/web-service-gin/src/domain"
	repository_interface "example/web-service-gin/src/repository/interface"

	"github.com/gin-gonic/gin"
)

type AlbumGetUsecase struct {
	albumRepo repository_interface.AlbumRepository
}

func NewAlbumGetUsecase(
	albumRepo repository_interface.AlbumRepository,
) *AlbumGetUsecase {
	return &AlbumGetUsecase{
		albumRepo: albumRepo,
	}
}

func (usecase *AlbumGetUsecase) Execute(
	ctx *gin.Context,
	id string,
) (*domain.Album, error) {
	album, err := usecase.albumRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return album, nil
}
