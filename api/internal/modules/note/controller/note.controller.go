package controller

import (
	"net/http"
	"strconv"

	"github.co/golang-programming/restaurant/api/note/dto"
	"github.co/golang-programming/restaurant/api/note/service"
	"github.com/gin-gonic/gin"
)

func CreateNote(ctx *gin.Context) {
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.CreateNoteInput)

	note, err := service.CreateNote(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, note)
}

func GetNoteByID(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	note, err := service.GetNoteByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, note)
}

func UpdateNote(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	val, _ := ctx.Get("validatedInput")
	input := val.(*dto.UpdateNoteInput)

	note, err := service.UpdateNote(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, note)
}

func DeleteNote(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	err := service.DeleteNote(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func ListNotes(ctx *gin.Context) {
	notes, err := service.ListNotes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notes)
}
