package websaystest

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.Run("localhost:8080")
}