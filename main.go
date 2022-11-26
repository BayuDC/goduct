package main

import (
	"goduct/controllers"
	"goduct/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/api/products", controllers.ProductIndex)
	r.GET("/api/products/:id", controllers.ProductShow)
	r.POST("/api/products", controllers.ProductStore)
	r.PUT("/api/products/:id", controllers.ProductUpdate)
	r.DELETE("/api/products/:id", controllers.ProductDestroy)

	r.Run()
}
