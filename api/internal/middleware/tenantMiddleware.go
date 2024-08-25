package middleware

import (
	"net/http"

	"github.co/golang-programming/restaurant/api/pkg/database"
	"github.com/gin-gonic/gin"
)

// TenantMiddleware extracts the X-Tenant-ID header and sets the appropriate DB context
func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.GetHeader("X-Tenant-ID")
		if tenantID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header missing"})
			c.Abort()
			return
		}

		err := database.SetDB(tenantID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid tenant ID"})
			c.Abort()
			return
		}

		// Continue to the next handler
		c.Next()
	}
}
