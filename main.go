package main

import (
	"encoding/json"
	"log"

	"github.com/DumbBoi/websays-test/models"
	"github.com/gin-gonic/gin"
)

func createArticle(c *gin.Context) {
	var (
		request models.ArticleRequest
		err     error
	)

	err = c.BindJSON(&request)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, err.Error())
		return
	}
	models.CreateArticle(request)
}

func ReadAritcle(c *gin.Context) {
	var (
		err     error
		request struct {
			key string
		}
	)
	err = c.BindJSON(&request)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, err.Error())
		return
	}
	art := models.ReadAritcle(request.key)
	output, err := json.Marshal(art)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, output)
}

func main() {
	router := gin.Default()
	router.POST("/article", createArticle)
	router.GET("/article", ReadAritcle)

	router.POST("/category", createArticle)
	router.GET("/category", ReadAritcle)
	router.DELETE("/category", createArticle)
	router.PUT("/category", createArticle)

	router.Run("localhost:8080")
}
