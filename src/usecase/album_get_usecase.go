package usecase

import (
	"example/web-service-gin/src/domain"
	repository_interface "example/web-service-gin/src/repository/interface"

	"github.com/gin-gonic/gin"
)

type GetAlbumUsecase struct {
	albumRepo repository_interface.AlbumRepository
}

func NewGetAlbumUsecase(
	albumRepo repository_interface.AlbumRepository,
) *GetAlbumUsecase {
	return &GetAlbumUsecase{
		albumRepo: albumRepo,
	}
}

func (usecase *GetAlbumUsecase) Execute(
	ctx *gin.Context,
	id string,
) (*domain.Album, error) {
	album, err := usecase.albumRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return album, nil
}
