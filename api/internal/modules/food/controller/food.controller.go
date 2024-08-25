package controller

import (
	"net/http"
	"strconv"

	"github.co/golang-programming/restaurant/api/internal/modules/food/dto"
	"github.co/golang-programming/restaurant/api/internal/modules/food/service"
	"github.com/gin-gonic/gin"
)

func CreateFood(ctx *gin.Context) {
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.CreateFoodInput)

	food, err := service.CreateFood(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, food)
}

func GetFoodByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	food, err := service.GetFoodByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, food)
}

func UpdateFood(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.UpdateFoodInput)

	food, err := service.UpdateFood(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, food)
}

func ChangeFoodStatus(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.ChangeFoodStatusInput)

	err := service.ChangeFoodStatus(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

func DeleteFood(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	err := service.DeleteFood(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func ListFoods(ctx *gin.Context) {
	foods, err := service.ListFoods()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, foods)
}

func Update(ctx *gin.Context) {
	foods, err := service.ListFoods()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, foods)
}
