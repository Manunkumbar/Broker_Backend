package main

import (
	"broker-backend/controllers"
	"broker-backend/database"
	"broker-backend/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database.ConnectDB()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/refresh", controllers.RefreshToken)

	private := r.Group("/")
	private.GET("/holdings", middleware.JWTMiddleware(), controllers.GetHoldings)
	private.GET("/orderbook", middleware.JWTMiddleware(), controllers.GetOrderbook)
	private.GET("/positions", middleware.JWTMiddleware(), controllers.GetPositions)

	r.Run(":6001") // Run the server on port 6001
}
