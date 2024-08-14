package middleware

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate *validator.Validate

// Initialize the validator
func init() {
	validate = validator.New()
}

// InputValidator is a middleware that validates the input based on the provided struct type
func InputValidator(inputType interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new instance of the input struct type
		input := createInstance(inputType)

		// Bind the JSON input to the new instance
		if err := c.ShouldBindJSON(input); err != nil {
			// Check if the error is due to an unknown field
			if strings.Contains(err.Error(), "json: unknown field") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input. Please use lowercase keys."})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
			}
			c.Abort()
			return
		}

		// Validate the input using the validator
		if err := validate.Struct(input); err != nil {
			// Collect all validation errors
			var errors []string
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.StructNamespace()+": "+err.ActualTag())
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			c.Abort()
			return
		}

		// Save the validated input to the context
		c.Set("validatedInput", input)

		// Continue to the next middleware/handler
		c.Next()
	}
}

// Create a new instance of the input struct type
func createInstance(inputType interface{}) interface{} {
	return reflect.New(reflect.TypeOf(inputType).Elem()).Interface()
}
