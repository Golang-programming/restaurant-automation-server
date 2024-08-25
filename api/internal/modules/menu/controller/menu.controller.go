package controller

import (
	"net/http"
	"strconv"

	"github.co/golang-programming/restaurant/api/internal/modules/menu/dto"
	"github.co/golang-programming/restaurant/api/internal/modules/menu/service"
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
	foodID, _ := strconv.Atoi(ctx.Param("food_id"))

	err := service.AddFoodToMenu(uint(id), uint(foodID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

func RemoveFoodFromMenu(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	foodID, _ := strconv.Atoi(ctx.Param("food_id"))

	err := service.RemoveFoodFromMenu(uint(id), uint(foodID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

func AddDealToMenu(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	dealID, _ := strconv.Atoi(ctx.Param("deal_id"))

	err := service.AddDealToMenu(uint(id), uint(dealID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

func RemoveDealFromMenu(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	dealID, _ := strconv.Atoi(ctx.Param("deal_id"))

	err := service.RemoveDealFromMenu(uint(id), uint(dealID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
