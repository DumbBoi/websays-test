package websaystest

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/article", createArticle)
	router.Run("localhost:8080")
}
