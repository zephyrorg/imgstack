package http-server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func api-serve() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Println("listen·and·serve·on·0.0.0.0:52414")
	router.Run(":52414") // listen and serve on 0.0.0.0:8080
}
