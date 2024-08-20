package bill

import (
	"github.co/golang-programming/restaurant/api/bill/controller"
	"github.co/golang-programming/restaurant/api/bill/dto"
	"github.co/golang-programming/restaurant/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/bill")

	groupRouter.POST("/", middleware.InputValidator(&dto.CreateBillInput{}), controller.CreateBill)
	groupRouter.GET("/:id", controller.GetBillByID)
	groupRouter.PUT("/:id", middleware.InputValidator(&dto.UpdateBillInput{}), controller.UpdateBill)
	groupRouter.PUT("/:id/mark-as-paid", middleware.InputValidator(&dto.MarkBillPaidInput{}), controller.MarkBillPaidInput)
	groupRouter.DELETE("/:id", controller.DeleteBill)
	groupRouter.GET("/", controller.ListBills)
}
