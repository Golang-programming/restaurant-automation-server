package main

import (
	"github.co/golang-programming/restaurant/src/food"
	"github.co/golang-programming/restaurant/src/invoice"
	"github.co/golang-programming/restaurant/src/menu"
	"github.co/golang-programming/restaurant/src/note"
	"github.co/golang-programming/restaurant/src/order"
	"github.co/golang-programming/restaurant/src/table"
	"github.co/golang-programming/restaurant/src/user"
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
