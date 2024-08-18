package controller

import (
	"net/http"
	"strconv"

	"github.co/golang-programming/restaurant/api/staff/dto"
	"github.co/golang-programming/restaurant/api/staff/service"
	"github.com/gin-gonic/gin"
)

func CreateStaff(ctx *gin.Context) {
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.CreateStaffInput)

	Staff, err := service.CreateStaff(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, Staff)
}

func UpdateStaff(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.UpdateStaffInput)

	Staff, err := service.UpdateStaff(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Staff)
}

func GetAllStaffs(ctx *gin.Context) {
	staffs, err := service.GetAllStaffs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, staffs)
}

func GetStaffByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	Staff, err := service.GetStaffByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Staff)
}

func DeleteStaff(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	err := service.DeleteStaff(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
