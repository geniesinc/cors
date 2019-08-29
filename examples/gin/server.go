package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/axiomzen/cors/v2/wrapper/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	router.Run(":8080")
}
