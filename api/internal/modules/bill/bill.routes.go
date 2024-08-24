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
	groupRouter.PUT("/:id/mark-as-paid", middleware.InputValidator(&dto.MarkBillPaidInput{}), controller.MarkBillPaid)
	groupRouter.PUT("/:id/refresh", controller.RefreshBill)
	groupRouter.PUT("/:id/discount", middleware.InputValidator(&dto.UpdateDiscountInput{}), controller.UpdateDiscount)
	groupRouter.DELETE("/:id", controller.DeleteBill)
	groupRouter.GET("/", controller.ListBills)
}