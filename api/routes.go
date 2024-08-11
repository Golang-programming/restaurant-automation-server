package main

import (
	"github.co/golang-programming/restaurant/api/food"
	"github.co/golang-programming/restaurant/api/invoice"
	"github.co/golang-programming/restaurant/api/menu"
	"github.co/golang-programming/restaurant/api/note"
	"github.co/golang-programming/restaurant/api/order"
	"github.co/golang-programming/restaurant/api/table"
	"github.co/golang-programming/restaurant/api/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	user.RegisterRoutes(router)
	note.RegisterRoutes(router)
	food.RegisterRoutes(router)
	invoice.RegisterRoutes(router)
	menu.RegisterRoutes(router)
	order.RegisterRoutes(router)
	table.RegisterRoutes(router)
}
