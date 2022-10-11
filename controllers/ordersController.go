package controllers

import (
	"assignment-2/database"
	"assignment-2/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var OrderDatas = []models.Orders{}

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Orders{}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newOrder := models.Orders{
		OrderedAt:    order.OrderedAt,
		CustomerName: order.CustomerName,
		Items:        order.Items,
	}

	if err := db.Create(&newOrder).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"order": newOrder,
	})
}

func GetOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Orders{}

	err := db.First(&order, "order_id = ?", order.OrderID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Data Not Found")
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

func UpdateOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Orders{}

	err := db.Model(&order).Where("order_id = ?", order.OrderID).Updates(models.Orders{OrderID: order.OrderID}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Orders{}

	err := db.Where("order_id= ?", order.OrderID).Delete(&order).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order Data Successfully Deleted",
	})
}