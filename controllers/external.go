package controllers

import (
	"broker-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CallUpstream(c *gin.Context) {
	url := "https://jsonplaceholder.typicode.com/todos/1" // Example upstream

	data, err := utils.GetWithCircuitBreaker(url)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", data)
}
