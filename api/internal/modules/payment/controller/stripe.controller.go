package controller

import (
	"net/http"
	"strconv"

	"github.co/golang-programming/restaurant/api/payment/service"
	"github.com/gin-gonic/gin"
)

func CreatePaymentIntent(ctx *gin.Context) {

	id, _ := strconv.ParseUint(ctx.Param("customer_id"), 10, 32)

	clientSecret, err := service.CreatePaymentIntent(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"clientSecret": clientSecret,
	})
}
