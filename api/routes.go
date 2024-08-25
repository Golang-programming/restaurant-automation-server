package main

import (
	"github.co/golang-programming/restaurant/api/internal/modules/auth"
	"github.co/golang-programming/restaurant/api/internal/modules/bill"
	"github.co/golang-programming/restaurant/api/internal/modules/customer"
	"github.co/golang-programming/restaurant/api/internal/modules/menu"
	"github.co/golang-programming/restaurant/api/internal/modules/order"
	"github.co/golang-programming/restaurant/api/internal/modules/payment"
	staff "github.co/golang-programming/restaurant/api/internal/modules/staff"
	"github.co/golang-programming/restaurant/api/internal/modules/table"
	"github.co/golang-programming/restaurant/api/internal/websocket"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all the routes for the application.
func RegisterRoutes(router *gin.RouterGroup) {
	// Basic health check route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Register HTTP routes
	menu.RegisterRoutes(router)
	customer.RegisterRoutes(router)
	order.RegisterRoutes(router)
	table.RegisterRoutes(router)
	auth.RegisterRoutes(router)
	staff.RegisterRoutes(router)
	payment.RegisterRoutes(router)

	// Register WebSocket routes
	registerWebSocketRoutes(router)
}

// registerWebSocketRoutes registers WebSocket routes for different modules.
func registerWebSocketRoutes(router *gin.RouterGroup) {

	orderWs := bill.NewBillWebSocket(hub, natsClient)

	//
	router.GET("/ws/order", func(c *gin.Context) {
		websocket.HandleWebSocket(c, orderWs)
	})

	// // Customer WebSocket route
	// customerWs := customer.NewCustomerWebSocket(hub, natsClient)
	// router.GET("/ws/customer", func(c *gin.Context) {
	// 	websocket.HandleWebSocket(c, customerWs)
	// })

	// // Menu WebSocket route
	// menuWs := menu.NewMenuWebSocket(hub, natsClient)
	// router.GET("/ws/menu", func(c *gin.Context) {
	// 	websocket.HandleWebSocket(c, menuWs)
	// })
}
