package controller

import (
	"net/http"
	"strconv"

	"github.co/golang-programming/restaurant/api/order/dto"
	"github.co/golang-programming/restaurant/api/order/service"
	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.CreateOrderInput)

	order, err := service.CreateOrder(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

func GetOrderByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	order, err := service.GetOrderByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func UpdateOrder(ctx *gin.Context) {
	val, _ := ctx.Get("validatedInput")
	id, _ := strconv.Atoi(ctx.Param("id"))
	input := val.(*dto.UpdateOrderInput)

	order, err := service.UpdateOrder(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func DeleteOrder(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := service.DeleteOrder(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func ListOrders(ctx *gin.Context) {
	orders, err := service.ListOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

// Order Items
func AddOrderItem(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.AddOrderItemInput)

	orderItem, err := service.AddOrderItem(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderItem)
}

func UpdateOrderItem(ctx *gin.Context) {
	orderID, _ := strconv.Atoi(ctx.Param("order_id"))
	itemID, _ := strconv.Atoi(ctx.Param("item_id"))
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.UpdateOrderItemInput)

	orderItem, err := service.UpdateOrderItem(uint(orderID), uint(itemID), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orderItem)
}

func RemoveOrderItem(ctx *gin.Context) {
	orderID, _ := strconv.Atoi(ctx.Param("order_id"))
	itemID, _ := strconv.Atoi(ctx.Param("item_id"))

	if err := service.RemoveOrderItem(uint(orderID), uint(itemID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
