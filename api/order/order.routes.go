package order

import (
	"github.co/golang-programming/restaurant/api/middleware"
	"github.co/golang-programming/restaurant/api/order/controller"
	"github.co/golang-programming/restaurant/api/order/dto"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/orders")

	groupRouter.POST("/", middleware.InputValidator(&dto.CreateOrderInput{}), controller.CreateOrder)
	groupRouter.GET("/:order_id", controller.GetOrderByID)
	groupRouter.PUT("/:order_id", middleware.InputValidator(&dto.UpdateOrderStatusInput{}), controller.UpdateOrderStatus)
	groupRouter.DELETE("/:order_id", controller.DeleteOrder)
	groupRouter.GET("/", controller.ListOrders)

	// Routes for managing order items
	groupRouter.POST("/:order_id/item", middleware.InputValidator(&dto.AddOrderItemInput{}), controller.AddOrderItem)
	groupRouter.PUT("/:order_id/item/:item_id", middleware.InputValidator(&dto.UpdateOrderItemStatusInput{}), controller.UpdateOrderItemStatus)
	groupRouter.DELETE("/:order_id/item/:item_id", controller.RemoveOrderItem)
}
