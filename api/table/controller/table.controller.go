package controller

import (
	"net/http"
	"strconv"

	"github.co/golang-programming/restaurant/api/table/dto"
	"github.co/golang-programming/restaurant/api/table/service"
	"github.com/gin-gonic/gin"
)

func CreateTable(ctx *gin.Context) {
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.CreateTableInput)

	table, err := service.CreateTable(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, table)
}

func GetTableByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	table, err := service.GetTableByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, table)
}

func UpdateTable(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.UpdateTableInput)

	table, err := service.UpdateTable(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, table)
}

func DeleteTable(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	err := service.DeleteTable(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func ListTables(ctx *gin.Context) {
	tables, err := service.ListTables()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tables)
}
