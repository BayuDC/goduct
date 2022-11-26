package controllers

import (
	"goduct/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductIndex(c *gin.Context) {
	var products []models.Product

	models.Db.Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
func ProductShow(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.Db.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Product not found",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}
func ProductStore(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	models.Db.Create(&product)
	c.JSON(http.StatusCreated, gin.H{
		"message": product,
	})

}
func ProductUpdate(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if models.Db.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}
func ProductDestroy(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if models.Db.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
