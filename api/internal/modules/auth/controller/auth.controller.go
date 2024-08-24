package controller

import (
	"net/http"

	"github.co/golang-programming/restaurant/api/auth/dto"
	"github.co/golang-programming/restaurant/api/auth/service"
	"github.com/gin-gonic/gin"
)

func SendOTP(ctx *gin.Context) {
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.SendOTPInput)

	err := service.SendOTP(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "opt sent to youru phone number"})
}

func ValidateOTP(ctx *gin.Context) {
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.ValidateOTPInput)

	err := service.ValidateOTP(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, user, err := service.LoginStaff(input.PhoneNumber)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken, "user": user})
}
