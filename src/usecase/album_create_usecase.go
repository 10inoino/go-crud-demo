package usecase

import (
	"example/web-service-gin/src/domain"
	repository_interface "example/web-service-gin/src/repository/interface"

	"github.com/gin-gonic/gin"
)

type AlbumCreateUsecase struct {
	albumRepo repository_interface.AlbumRepository
}

func NewAlbumCreateUsecase(
	albumRepo repository_interface.AlbumRepository,
) *AlbumCreateUsecase {
	return &AlbumCreateUsecase{
		albumRepo: albumRepo,
	}
}

func (usecase *AlbumCreateUsecase) Execute(
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
