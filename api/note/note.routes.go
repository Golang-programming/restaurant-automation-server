package note

import (
	"github.co/golang-programming/restaurant/api/middleware"
	"github.co/golang-programming/restaurant/api/note/controller"
	"github.co/golang-programming/restaurant/api/note/dto"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	groupRouter := router.Group("/note")

	groupRouter.POST("/", middleware.InputValidator(&dto.CreateNoteInput{}), controller.CreateNote)
	groupRouter.GET("/:id", controller.GetNoteByID)
	groupRouter.PUT("/:id", middleware.InputValidator(&dto.UpdateNoteInput{}), controller.UpdateNote)
	groupRouter.DELETE("/:id", controller.DeleteNote)
	groupRouter.GET("/", controller.ListNotes)
}
