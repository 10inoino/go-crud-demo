package usecase

import (
	repository_interface "example/web-service-gin/src/repository/interface"

	"github.com/gin-gonic/gin"
)

type DeleteAlbumUsecase struct {
	albumRepo repository_interface.AlbumRepository
}

func NewDeleteAlbumUsecase(
	albumRepo repository_interface.AlbumRepository,
) *DeleteAlbumUsecase {
	return &DeleteAlbumUsecase{
		albumRepo: albumRepo,
	}
}

func (usecase *DeleteAlbumUsecase) Execute(
	ctx *gin.Context,
	id string,
) error {
	_, err := usecase.albumRepo.FindById(ctx, id)
	if err != nil {
		return err
	}
	err = usecase.albumRepo.DeleteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
