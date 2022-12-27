package main

import (
	"database/sql"
	"errors"
	"example/web-service-gin/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  int    `json:"price"`
}

func NewAlbum(id string, title string, artist string, price int) (*album, error) {
	return &album{
		ID:     id,
		Title:  title,
		Artist: artist,
		Price:  price,
	}, nil
}

type AlbumRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{
		db: db,
	}
}

func (repo *AlbumRepository) Save(ctx *gin.Context, album album) error {
	saveTarget := &models.Album{
		ID:     album.ID,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}

	err := saveTarget.Insert(ctx, repo.db, boil.Infer())
	return err
}

func (repo *AlbumRepository) FindAll(ctx *gin.Context) (*[]album, error) {
	m, err := models.Albums().All(ctx, repo.db)
	if err != nil {
		return nil, errors.New("failed get albums")
	}

	result := make([]album, len(m))
	// TODO:関数化したい
	for i, a := range m {
		album, _ := NewAlbum(a.ID, a.Title, a.Artist, a.Price)
		result[i] = *album
	}
	return &result, nil
}

func (repo *AlbumRepository) FindById(ctx *gin.Context, id string) (*album, error) {
	m, err := models.Albums(
		qm.Where("id=?", id),
	).One(ctx, repo.db)
	if err != nil {
		return nil, errors.New("failed find album")
	}
	album, _ := NewAlbum(m.ID, m.Title, m.Artist, m.Price)
	return album, nil
}

func generateDB() (*sql.DB, error) {
	return sql.Open(
		"postgres",
		"host=postgresql dbname=go-demo user=go-demo password=password sslmode=disable")
}

func getAlbums(c *gin.Context) {
	db, dbErr := generateDB()
	if dbErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed database connection"})
		return
	}
	albumRepo := NewAlbumRepository(db)
	album, err := albumRepo.FindAll(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func postAlbums(c *gin.Context) {
	db, dbErr := generateDB()
	if dbErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed database connection"})
		return
	}
	albumRepo := NewAlbumRepository(db)
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed bind json"})
		return
	}
	err := albumRepo.Save(c, newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusCreated, "OK")
}

func getAlbumByID(c *gin.Context) {
	db, dbErr := generateDB()
	if dbErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed database connection"})
		return
	}
	albumRepo := NewAlbumRepository(db)
	id := c.Param("id")
	album, err := albumRepo.FindById(c, id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed find data"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
