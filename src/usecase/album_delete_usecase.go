package usecase

import (
	repository_interface "example/web-service-gin/src/repository/interface"

	"github.com/gin-gonic/gin"
)

type AlbumDeleteUsecase struct {
	albumRepo repository_interface.AlbumRepository
}

func NewAlbumDeleteUsecase(
	albumRepo repository_interface.AlbumRepository,
) *AlbumDeleteUsecase {
	return &AlbumDeleteUsecase{
		albumRepo: albumRepo,
	}
}

func (usecase *AlbumDeleteUsecase) Execute(
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
