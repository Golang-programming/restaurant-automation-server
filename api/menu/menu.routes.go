package menu

import (
	"github.co/golang-programming/restaurant/api/menu/controller"
	"github.co/golang-programming/restaurant/api/menu/dto"
	"github.co/golang-programming/restaurant/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/menu")

	groupRouter.POST("/", middleware.InputValidator(&dto.CreateMenuInput{}), controller.CreateMenu)
	groupRouter.GET("/:id", controller.GetMenuByID)
	groupRouter.PUT("/:id", middleware.InputValidator(&dto.UpdateMenuInput{}), controller.UpdateMenu)
	groupRouter.DELETE("/:id", controller.DeleteMenu)
	groupRouter.GET("/", controller.ListMenus)

	// Routes for adding/removing foods and deals
	groupRouter.POST("/:id/food", middleware.InputValidator(&dto.AddFoodToMenuInput{}), controller.AddFoodToMenu)
	groupRouter.DELETE("/:id/food", middleware.InputValidator(&dto.RemoveFoodFromMenuInput{}), controller.RemoveFoodFromMenu)
	groupRouter.POST("/:id/deal", middleware.InputValidator(&dto.AddDealToMenuInput{}), controller.AddDealToMenu)
	groupRouter.DELETE("/:id/deal", middleware.InputValidator(&dto.RemoveDealFromMenuInput{}), controller.RemoveDealFromMenu)
}
