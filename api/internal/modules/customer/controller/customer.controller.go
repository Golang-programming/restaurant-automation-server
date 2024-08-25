package controller

import (
	"net/http"
	"strconv"

	"github.co/golang-programming/restaurant/api/internal/modules/customer/dto"
	"github.co/golang-programming/restaurant/api/internal/modules/customer/service"
	"github.com/gin-gonic/gin"
)

func CreateCustomer(ctx *gin.Context) {
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.CreateCustomerInput)

	customer, err := service.CreateCustomer(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, customer)
}

func GetCustomerByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	customer, err := service.GetCustomerByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func UpdateCustomer(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.UpdateCustomerInput)

	customer, err := service.UpdateCustomer(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func DeleteCustomer(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	err := service.DeleteCustomer(uint(id))
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

func DeactivateCustomer(ctx *gin.Context) {
	val, _ := ctx.Get("currentCustomerID")
	id, _ := strconv.ParseUint(val.(string), 10, 32)

	err := service.DeactivateCustomer(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Customer deactivated successfully"})
}
