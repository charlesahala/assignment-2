package controllers

import (
	"assignment-2/database"
	"assignment-2/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

var OrderDatas = []models.Order{}

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Please Fill The Required Data",
		})
		return
	}

	newOrder := models.Order{
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

	ctx.JSON(http.StatusCreated, newOrder)
}

func GetOrder(ctx *gin.Context) {
	db := database.GetDB()
	orders := []models.Order{}

	if err := ctx.ShouldBind(&orders); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := db.Model(&models.Order{}).Preload("Items").Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	ordersData := make([]models.Order, len(orders))

	for i := range orders {
		itemsData := make([]models.Item, len(orders[i].Items))

		for j := range orders[i].Items {
			itemsData[j] = models.Item{
				Description: orders[i].Items[j].Description,
				ItemID:      orders[i].Items[j].ItemID,
				Quantity:    orders[i].Items[j].Quantity,
			}

			ordersData[i] = models.Order{
				CustomerName: orders[i].CustomerName,
				Items:        itemsData,
				OrderID:      orders[i].OrderID,
				OrderedAt:    orders[i].OrderedAt,
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order data": ordersData,
	})
}

func UpdateOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}
	item := models.Item{}

	if err := db.Where("order_id = ?", ctx.Param("orderID")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})

		return
	}

	db.Unscoped().Where("order_id = ?", order.OrderID).Delete(item)

	if err := db.Unscoped().Where("order_id = ?", order.OrderID).Delete(item).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := ctx.ShouldBind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := db.Save(order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	updatedItem := make([]models.Item, len(order.Items))

	for i := range order.Items {
		updatedItem[i] = models.Item{
			Description: order.Items[i].Description,
			ItemCode:    order.Items[i].ItemCode,
			Quantity:    order.Items[i].Quantity,
		}
	}

	updatedOrder := models.Order{
		CustomerName: order.CustomerName,
		Items:        updatedItem,
		OrderID:      order.OrderID,
		OrderedAt:    order.OrderedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order update": updatedOrder,
	})

	// err := db.Model(&order).Where("order_id = ?", order.OrderID).Updates(models.Order{OrderID: order.OrderID}).Error

	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error":   "Bad Request",
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"data": order,
	// })
}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}

	if err := db.Where("order_id = ?", ctx.Param("orderID")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := db.Select(clause.Associations).Delete(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order Successfully Deleted",
	})
}
