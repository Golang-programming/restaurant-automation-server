package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	LoadEnv()
	PORT := os.Getenv("PORT")

	app := gin.New()

	v1 := app.Group("/v1")
	RegisterRoutes(v1)

	fmt.Println("App running on port: ", PORT)
	app.Run(":" + PORT)
}
