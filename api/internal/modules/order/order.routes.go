package order

import (
	"github.co/golang-programming/restaurant/api/internal/middleware"
	"github.co/golang-programming/restaurant/api/internal/modules/order/controller"
	"github.co/golang-programming/restaurant/api/internal/modules/order/dto"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/orders")

	groupRouter.POST("/", middleware.InputValidator(&dto.CreateOrderInput{}), controller.CreateOrder)
	groupRouter.GET("/:id", controller.GetOrderByID)
	groupRouter.PUT("/status/:id", middleware.InputValidator(&dto.UpdateOrderStatusInput{}), controller.UpdateOrderStatus)
	groupRouter.DELETE("/:id", controller.DeleteOrder)
	groupRouter.GET("/", controller.ListOrders)

	// Routes for managing order items
	groupRouter.POST("/:id/item", middleware.InputValidator(&dto.AddOrderItemInput{}), controller.AddOrderItem)
	groupRouter.PUT("/:id/item/:item_id", middleware.InputValidator(&dto.UpdateOrderItemStatusInput{}), controller.UpdateOrderItemStatus)
	groupRouter.DELETE("/:id/item/:item_id", controller.RemoveOrderItem)
}
