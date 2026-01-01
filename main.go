package main

import (
	"go-crud/src/config"
	"go-crud/src/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	config.InitEnv()
	config.Connection()

	config.DB.AutoMigrate(&entity.Product{})
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "get pong",
		})
	})
	router.Run()
}
