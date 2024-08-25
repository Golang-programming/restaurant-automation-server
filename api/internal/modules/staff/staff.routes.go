package staff

import (
	"github.co/golang-programming/restaurant/api/internal/middleware"
	"github.co/golang-programming/restaurant/api/internal/modules/staff/controller"
	"github.co/golang-programming/restaurant/api/internal/modules/staff/dto"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/staff")

	groupRouter.POST("/", middleware.InputValidator(&dto.CreateStaffInput{}), controller.CreateStaff)
	groupRouter.PUT("/", middleware.InputValidator(&dto.UpdateStaffInput{}), controller.UpdateStaff)
	groupRouter.GET("/", controller.GetAllStaffs)      // Route to get all staffs
	groupRouter.GET("/:id", controller.GetStaffByID)   // Route to get a Staff by ID
	groupRouter.DELETE("/:id", controller.DeleteStaff) // Route to delete a Staff by ID
}
