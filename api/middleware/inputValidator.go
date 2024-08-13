package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func InputValidator(obj interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if err := ctx.ShouldBindJSON(obj); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
			ctx.Abort()
			return
		}

		if err := validate.Struct(obj); err != nil {
			errors := make([]string, 0)
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.StructField()+" "+err.ActualTag())
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": errors})
			ctx.Abort()
			return
		}

		ctx.Set("validatedBody", obj)
		ctx.Next()
	}
}
