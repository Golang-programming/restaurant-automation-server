package controller

import (
	"net/http"
	"strconv"

	"github.co/golang-programming/restaurant/api/menu/dto"
	"github.co/golang-programming/restaurant/api/menu/service"
	"github.com/gin-gonic/gin"
)

func CreateMenu(ctx *gin.Context) {
	var input dto.CreateMenuInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	menu, err := service.CreateMenu(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, menu)
}

func GetMenuByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}
	menu, err := service.GetMenuByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}
	ctx.JSON(http.StatusOK, menu)
}

func UpdateMenu(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}
	var input dto.UpdateMenuInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	menu, err := service.UpdateMenu(uint(id), &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, menu)
}

func DeleteMenu(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
		return
	}
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
