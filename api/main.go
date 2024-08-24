package main

import (
	"fmt"
	"os"

	"github.co/golang-programming/restaurant/api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	LoadEnv()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9001"
	}

	app := gin.New()

	v1 := app.Group("/v1")

	v1.Use(middleware.TenantMiddleware())

	RegisterRoutes(v1)

	fmt.Println("App running on port: ", PORT)
	app.Run(":" + PORT)
}
