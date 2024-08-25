package customer

import (
	"github.co/golang-programming/restaurant/api/internal/middleware"
	"github.co/golang-programming/restaurant/api/internal/modules/customer/controller"
	"github.co/golang-programming/restaurant/api/internal/modules/customer/dto"
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
