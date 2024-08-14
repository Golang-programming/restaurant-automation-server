package middleware

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func InputValidator(inputType interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		input := createInstance(inputType)

		if err := c.ShouldBindJSON(input); err != nil {
			if strings.Contains(err.Error(), "json: unknown field") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input. Please use lowercase keys."})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			c.Abort()
			return
		}

		if err := validate.Struct(input); err != nil {
			var errors []string
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.StructNamespace()+": "+err.ActualTag())
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			c.Abort()
			return
		}

		c.Set("validatedInput", input)

		c.Next()
	}
}

// Create a new instance of the input struct type
func createInstance(inputType interface{}) interface{} {
	return reflect.New(reflect.TypeOf(inputType).Elem()).Interface()
}
