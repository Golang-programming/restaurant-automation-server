package controller

import (
	"net/http"

	"github.co/golang-programming/restaurant/api/customer/dto"
	"github.co/golang-programming/restaurant/api/customer/service"
	"github.com/gin-gonic/gin"
)

func CreateCustomer(ctx *gin.Context) {
	var input dto.CreateCustomerInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err := service.CreateCustomer(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, customer)
}

func GetCustomerByID(ctx *gin.Context) {
	id := ctx.Param("id")
	customer, err := service.GetCustomerByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

func UpdateCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	var input dto.UpdateCustomerInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err := service.UpdateCustomer(id, &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

func DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	err := service.DeleteCustomer(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func ListCustomers(ctx *gin.Context) {
	customers, err := service.ListCustomers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customers)
}
