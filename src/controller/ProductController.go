package controller

import (
	"go-crud/src/config"
	"go-crud/src/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductIndex(c *gin.Context) {
	var products []entity.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func ProductPOST(c *gin.Context) {
	var product entity.Product
	c.Bind(&product)

	config.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{
		"message": "Product post success",
	})
}

func ProductPUT(c *gin.Context) {
	id := c.Param("id")
	var product entity.Product

	config.DB.First(&product, id)

	var updProduct entity.Product
	c.Bind(&updProduct)

	config.DB.Model(&product).Updates(entity.Product{
		Code:  updProduct.Code,
		Price: updProduct.Price,
	})
	c.JSON(http.StatusOK, gin.H{
		"message": "Product put success",
	})

}

// func ProductDelete(c *gin.Context) {

// }

// func ProductById(c *gin.Context) {

// }
