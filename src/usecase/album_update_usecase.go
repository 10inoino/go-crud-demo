package usecase

import (
	repository_interface "example/web-service-gin/src/repository/interface"

	"github.com/gin-gonic/gin"
)

type AlbumUpdateUsecase struct {
	albumRepo repository_interface.AlbumRepository
}

func NewAlbumUpdateUsecase(
	albumRepo repository_interface.AlbumRepository,
) *AlbumUpdateUsecase {
	return &AlbumUpdateUsecase{
		albumRepo: albumRepo,
	}
}

func (usecase *AlbumUpdateUsecase) Execute(
	ctx *gin.Context,
	id string,
	title string,
	artist string,
	price int,
) error {
	exist, err := usecase.albumRepo.FindById(ctx, id)
	if err != nil {
		return err
	}
	err = exist.Update(title, artist, price)
	if err != nil {
		return err
	}
	err = usecase.albumRepo.Update(ctx, *exist)
	if err != nil {
		return err
	}
	return nil
}
