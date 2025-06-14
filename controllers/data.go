package controllers

import (
	"broker-backend/database"
	"broker-backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHoldings(c *gin.Context) {
	email := c.MustGet("email").(string)
	fmt.Println(email)

	var user models.User
	err := database.DB.Get(&user, "SELECT id FROM users WHERE email=$1", email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	fmt.Println(user)

	var holdings []models.Holding
	err = database.DB.Select(&holdings, "SELECT user_id, stock_symbol, quantity, average_price FROM holdings WHERE user_id=$1", user.ID)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No holdings found"})
		return
	}

	c.JSON(http.StatusOK, holdings)
}

func GetOrderbook(c *gin.Context) {
	email := c.MustGet("email").(string)

	var user models.User
	err := database.DB.Get(&user, "SELECT id FROM users WHERE email=$1", email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	var orders []models.Order
	err = database.DB.Select(&orders, "SELECT user_id, stock_symbol, order_type, quantity, price, status FROM orderbook WHERE user_id=$1", user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No orders found"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetPositions(c *gin.Context) {
	email := c.MustGet("email").(string)

	var user models.User
	err := database.DB.Get(&user, "SELECT id FROM users WHERE email=$1", email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	var positions []models.Position
	err = database.DB.Select(&positions, "SELECT  user_id, stock_symbol, quantity, pnl FROM positions WHERE user_id=$1", user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No positions found"})
		return
	}

	c.JSON(http.StatusOK, positions)
}
