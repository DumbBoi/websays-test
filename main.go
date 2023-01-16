package main

import (
	"fmt"
	"log"

	"github.com/DumbBoi/websays-test/models"
	"github.com/gin-gonic/gin"
)

func createArticle(c *gin.Context) {
	var (
		request models.CreateArticleRequest
		err     error
	)

	err = c.BindJSON(&request)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, err.Error())
		return
	}
	art, _ := models.CreateArticle(request)
	fmt.Println(art.Id)
	fmt.Println(art.Title)
	fmt.Println(art.Categories)
}
func main() {
	router := gin.Default()
	router.POST("/article", createArticle)

	router.Run("localhost:8080")
}
