package main

import (
	"github.co/golang-programming/restaurant/api/auth"
	"github.co/golang-programming/restaurant/api/customer"
	"github.co/golang-programming/restaurant/api/menu"
	"github.co/golang-programming/restaurant/api/order"
	"github.co/golang-programming/restaurant/api/table"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	menu.RegisterRoutes(router)
	customer.RegisterRoutes(router)
	order.RegisterRoutes(router)
	table.RegisterRoutes(router)
	auth.RegisterRoutes(router)
}
