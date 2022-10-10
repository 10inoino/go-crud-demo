package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  string `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: "56.99"},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: "17.99"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: "39.99"},
}

var ddb = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "bind failed"})
		return
	}

	param1 := &dynamodb.PutItemInput{
		TableName: aws.String("gin-tutorial-table"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(newAlbum.ID),
			},
			"title": {
				S: aws.String(newAlbum.Title),
			},
			"artist": {
				S: aws.String(newAlbum.Artist),
			},
			"price": {
				N: aws.String(newAlbum.Price),
			},
		},
	}
	_, err := ddb.PutItem(param1)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
