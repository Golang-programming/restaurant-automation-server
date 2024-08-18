package auth

import (
	"github.co/golang-programming/restaurant/api/auth/controller"
	"github.co/golang-programming/restaurant/api/auth/dto"
	"github.co/golang-programming/restaurant/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/auth")
	groupRouter.POST("/send", middleware.InputValidator(&dto.SendOTPInput{}), controller.SendOTP)
	groupRouter.POST("/velidate", middleware.InputValidator(&dto.ValidateOTPInput{}), controller.ValidateOTP)
}
