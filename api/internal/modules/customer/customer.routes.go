package customer

import (
	"github.co/golang-programming/restaurant/api/customer/controller"
	"github.co/golang-programming/restaurant/api/customer/dto"
	"github.co/golang-programming/restaurant/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/customer")

	groupRouter.GET("/", controller.ListCustomers)
	groupRouter.GET("/:id", controller.GetCustomerByID)
	groupRouter.DELETE("/:id", controller.DeleteCustomer)
	groupRouter.POST("/", middleware.InputValidator(&dto.CreateCustomerInput{}), controller.CreateCustomer)
	groupRouter.PUT("/:id", middleware.InputValidator(&dto.UpdateCustomerInput{}), controller.UpdateCustomer)

	groupRouter.PUT("/deactivate", middleware.CustomerMiddleware(), controller.DeactivateCustomer)
}
