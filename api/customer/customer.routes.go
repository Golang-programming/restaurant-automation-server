package customer

import (
	"github.co/golang-programming/restaurant/api/customer/controller"
	"github.co/golang-programming/restaurant/api/customer/dto"
	"github.co/golang-programming/restaurant/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/customer")

	groupRouter.POST("/", middleware.InputValidator(&dto.CreateCustomerInput{}), controller.CreateCustomer)
	groupRouter.GET("/:id", controller.GetCustomerByID)
	groupRouter.PUT("/:id", middleware.InputValidator(&dto.UpdateCustomerInput{}), controller.UpdateCustomer)
	groupRouter.DELETE("/:id", controller.DeleteCustomer)
	groupRouter.GET("/", controller.ListCustomers)
}
