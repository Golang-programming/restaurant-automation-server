package auth

import (
	"github.co/golang-programming/restaurant/api/internal/middleware"
	"github.co/golang-programming/restaurant/api/internal/modules/auth/controller"
	"github.co/golang-programming/restaurant/api/internal/modules/auth/dto"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/auth")
	groupRouter.POST("/send", middleware.InputValidator(&dto.SendOTPInput{}), controller.SendOTP)
	groupRouter.POST("/velidate", middleware.InputValidator(&dto.ValidateOTPInput{}), controller.ValidateOTP)
}
