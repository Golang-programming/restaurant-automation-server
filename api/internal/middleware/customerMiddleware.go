package middleware

import (
	"net/http"

	"github.co/golang-programming/restaurant/api/pkg/utils/encryption"
	"github.com/gin-gonic/gin"
)

func CustomerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		customerCypherID := c.GetHeader("X-Customer-ID")
		if customerCypherID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header missing"})
			c.Abort()
			return
		}

		customerID, err := encryption.Decryptor(customerCypherID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Detected invalid customer"})
			c.Abort()
			return
		}

		c.Set("currentCustomerID", customerID)
		c.Next()
	}
}
