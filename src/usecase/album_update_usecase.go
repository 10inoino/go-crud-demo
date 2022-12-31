package usecase

import (
	repository_interface "example/web-service-gin/src/repository/interface"

	"github.com/gin-gonic/gin"
)

type UpdateAlbumUsecase struct {
	albumRepo repository_interface.AlbumRepository
}

func NewUpdateAlbumUsecase(
	albumRepo repository_interface.AlbumRepository,
) *UpdateAlbumUsecase {
	return &UpdateAlbumUsecase{
		albumRepo: albumRepo,
	}
}

func (usecase *UpdateAlbumUsecase) Execute(
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
