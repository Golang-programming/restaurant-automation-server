package food

import (
	"github.co/golang-programming/restaurant/api/food/controller"
	"github.co/golang-programming/restaurant/api/food/dto"
	"github.co/golang-programming/restaurant/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/food")

	groupRouter.POST("/", middleware.InputValidator(&dto.CreateFoodInput{}), controller.CreateFood)
	groupRouter.GET("/:id", controller.GetFoodByID)
	groupRouter.PUT("/:id", middleware.InputValidator(&dto.UpdateFoodInput{}), controller.UpdateFood)
	groupRouter.PUT("/:id/status", middleware.InputValidator(&dto.ChangeFoodStatusInput{}), controller.ChangeFoodStatus)
	groupRouter.DELETE("/:id", controller.DeleteFood)
	groupRouter.GET("/", controller.ListFoods)
}
