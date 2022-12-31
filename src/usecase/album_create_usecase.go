package usecase

import (
	"example/web-service-gin/src/domain"
	repository_interface "example/web-service-gin/src/repository/interface"

	"github.com/gin-gonic/gin"
)

type CreateAlbumUsecase struct {
	albumRepo repository_interface.AlbumRepository
}

func NewCreateAlbumUsecase(
	albumRepo repository_interface.AlbumRepository,
) *CreateAlbumUsecase {
	return &CreateAlbumUsecase{
		albumRepo: albumRepo,
	}
}

func (usecase *CreateAlbumUsecase) Execute(
	ctx *gin.Context,
	id string,
	title string,
	artist string,
	price int,
) error {
	exist, err := usecase.albumRepo.FindById(ctx, id)
	if exist != nil {
		return domain.NewDuplicateError("That album already exists.")
	}
	if !domain.IsNotFoundError(err) {
		return err
	}
	newAlbum, newAlbumErr := domain.NewAlbum(id, title, artist, price)
	if newAlbumErr != nil {
		return newAlbumErr
	}
	err = usecase.albumRepo.Save(ctx, *newAlbum)
	if err != nil {
		return err
	}
	return nil
}
