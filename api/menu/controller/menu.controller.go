package controller

import (
	"net/http"
	"strconv"

	"github.co/golang-programming/restaurant/api/menu/dto"
	"github.co/golang-programming/restaurant/api/menu/service"
	"github.com/gin-gonic/gin"
)

func CreateMenu(ctx *gin.Context) {
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.CreateMenuInput)

	menu, err := service.CreateMenu(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, menu)
}

func GetMenuByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	menu, err := service.GetMenuDetailsByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}

	ctx.JSON(http.StatusOK, menu)
}

func UpdateMenu(ctx *gin.Context) {
	val, _ := ctx.Get("validatedInput")
	id, _ := strconv.Atoi(ctx.Param("id"))
	input := val.(*dto.UpdateMenuInput)

	menu, err := service.UpdateMenu(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, menu)
}

func DeleteMenu(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := service.DeleteMenu(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func ListMenus(ctx *gin.Context) {
	menus, err := service.ListMenus()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, menus)
}

func AddFoodToMenu(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input dto.AddFoodToMenuInput
	_ = ctx.ShouldBindJSON(&input)

	menu, err := service.AddFoodToMenu(uint(id), &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, menu)
}

func RemoveFoodFromMenu(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input dto.RemoveFoodFromMenuInput
	_ = ctx.ShouldBindJSON(&input)

	menu, err := service.RemoveFoodFromMenu(uint(id), &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, menu)
}

func AddDealToMenu(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input dto.AddDealToMenuInput
	_ = ctx.ShouldBindJSON(&input)

	menu, err := service.AddDealToMenu(uint(id), &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, menu)
}

func RemoveDealFromMenu(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input dto.RemoveDealFromMenuInput
	_ = ctx.ShouldBindJSON(&input)

	menu, err := service.RemoveDealFromMenu(uint(id), &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, menu)
}
