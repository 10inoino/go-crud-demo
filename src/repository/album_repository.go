package repository

import (
	"database/sql"
	"errors"
	"example/web-service-gin/db/models"
	"example/web-service-gin/src/domain"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AlbumRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{
		db: db,
	}
}

func (repo *AlbumRepository) Save(ctx *gin.Context, album domain.Album) error {
	saveTarget := &models.Album{
		ID:     album.ID,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}

	err := saveTarget.Insert(ctx, repo.db, boil.Infer())
	return err
}

func (repo *AlbumRepository) FindAll(ctx *gin.Context) (*[]domain.Album, error) {
	m, err := models.Albums().All(ctx, repo.db)
	if err != nil {
		return nil, errors.New("failed get albums")
	}

	result := make([]domain.Album, len(m))
	// TODO:関数化したい
	for i, a := range m {
		album, _ := domain.NewAlbum(a.ID, a.Title, a.Artist, a.Price)
		result[i] = *album
	}
	return &result, nil
}

func (repo *AlbumRepository) FindById(ctx *gin.Context, id string) (*domain.Album, error) {
	m, err := models.Albums(
		qm.Where("id=?", id),
	).One(ctx, repo.db)
	if err != nil {
		return nil, errors.New("failed find album")
	}
	album, _ := domain.NewAlbum(m.ID, m.Title, m.Artist, m.Price)
	return album, nil
}