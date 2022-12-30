package repository

import (
	"database/sql"
	"example/web-service-gin/db/models"
	"example/web-service-gin/src/domain"
	repository_interface "example/web-service-gin/src/repository/interface"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AlbumPostgresRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) repository_interface.AlbumRepository {
	return &AlbumPostgresRepository{
		db: db,
	}
}

func (repo *AlbumPostgresRepository) Save(ctx *gin.Context, album domain.Album) error {
	saveTarget := &models.Album{
		ID:     album.ID,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}

	err := saveTarget.Insert(ctx, repo.db, boil.Infer())
	if err != nil {
		return domain.NewInvalidInputError("failed save album")
	}
	return nil
}

func (repo *AlbumPostgresRepository) FindAll(ctx *gin.Context) (*[]domain.Album, error) {
	m, err := models.Albums().All(ctx, repo.db)
	if err != nil {
		return nil, domain.NewInvalidInputError("failed find albums")
	}

	result := make([]domain.Album, len(m))
	for i, a := range m {
		album, _ := domain.NewAlbum(a.ID, a.Title, a.Artist, a.Price)
		result[i] = *album
	}
	return &result, nil
}

func (repo *AlbumPostgresRepository) FindById(ctx *gin.Context, id string) (*domain.Album, error) {
	m, err := models.FindAlbum(ctx, repo.db, id)
	if err == sql.ErrNoRows {
		return nil, domain.NewNotFoundError("album not found")
	}
	if err != nil {
		return nil, domain.NewInvalidInputError("failed find album")
	}
	album, _ := domain.NewAlbum(m.ID, m.Title, m.Artist, m.Price)
	return album, nil
}

func (repo *AlbumPostgresRepository) DeleteById(ctx *gin.Context, id string) error {
	m, err := models.FindAlbum(ctx, repo.db, id)
	if err == sql.ErrNoRows {
		return domain.NewNotFoundError("album not found")
	}
	if err != nil {
		return domain.NewInvalidInputError("failed find album")
	}
	_, delErr := m.Delete(ctx, repo.db)
	if delErr != nil {
		return err
	}
	return nil
}

func (repo *AlbumPostgresRepository) Update(ctx *gin.Context, album domain.Album) error {
	m, err := models.FindAlbum(ctx, repo.db, album.ID)
	if err == sql.ErrNoRows {
		return domain.NewNotFoundError("album not found")
	}
	if err != nil {
		return domain.NewInvalidInputError("failed find album")
	}
	m.ID = album.ID
	m.Title = album.Title
	m.Artist = album.Artist
	m.Price = album.Price
	_, delErr := m.Update(ctx, repo.db, boil.Infer())
	if delErr != nil {
		return domain.NewInvalidInputError("failed delete album")
	}
	return nil
}
