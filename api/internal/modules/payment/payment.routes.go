package payment

import (
	"github.co/golang-programming/restaurant/api/payment/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/stripe")

	groupRouter.POST("/:customer_id/payment-intent", controller.CreatePaymentIntent)
	/* groupRouter.POST("/", middleware.InputValidator(&dto.CreateFoodInput{}), controller.CreateFood)
	groupRouter.PUT("/:id", middleware.InputValidator(&dto.UpdateFoodInput{}), controller.UpdateFood)
	groupRouter.PUT("/:id/status", middleware.InputValidator(&dto.ChangeFoodStatusInput{}), controller.ChangeFoodStatus)
	groupRouter.DELETE("/:id", controller.DeleteFood)
	groupRouter.GET("/", controller.ListFoods) */
}
