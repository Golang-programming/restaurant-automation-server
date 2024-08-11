package main

import (
	"fmt"
	"os"

	"github.co/golang-programming/restaurant/api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	LoadEnv()
	PORT := os.Getenv("PORT")

	app := gin.New()

	v1 := app.Group("/v1")
	RegisterRoutes(v1)

	database.ConnectToDatabase()

	fmt.Println("App running on port: ", PORT)
	app.Run(":" + PORT)
}
