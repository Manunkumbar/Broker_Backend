package controllers

import (
	"broker-backend/database"
	"broker-backend/models"
	"broker-backend/utils"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	_, err = database.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := database.DB.QueryRow("SELECT id, email, password FROM users WHERE email=$1", req.Email).
		Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateJWT(user.Email, time.Minute*10)
	refresh, _ := utils.GenerateJWT(user.Email, time.Hour*24)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  token,
		"refresh_token": refresh,
	})
}

func RefreshToken(c *gin.Context) {
	var body struct {
		Token string `json:"token"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	email, err := utils.ValidateJWT(body.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	token, _ := utils.GenerateJWT(email, time.Minute*10)
	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
