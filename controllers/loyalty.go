package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tharmi/square-loyalty-backend/services"
	"github.com/tharmi/square-loyalty-backend/utils"
)

func Login(c *gin.Context) {
	utils.FakeLogin(c)
}

func EarnPoints(c *gin.Context) {
	var body struct {
		Points int64 `json:"points"`
	}

	if err := c.ShouldBindJSON(&body); err != nil || body.Points <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid points"})
		return
	}

	err := services.EarnPoints(body.Points)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to earn points"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Points earned successfully"})
}

func RedeemPoints(c *gin.Context) {
	var body struct {
		Points int64 `json:"points"`
	}

	if err := c.ShouldBindJSON(&body); err != nil || body.Points <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid points"})
		return
	}

	err := services.RedeemPoints(body.Points)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to redeem points"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Points redeemed successfully"})
}

func GetBalance(c *gin.Context) {
	balance, err := services.GetBalance()
	if err != nil {
		// Log the error to server console
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get balance", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

func GetHistory(c *gin.Context) {
	events, err := services.GetHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get history"})
		return
	}

	var history []map[string]interface{}
	for _, e := range events {
		eventType, _ := e["type"].(string)
		createdAt, _ := e["created_at"].(string)

		var points int64
		if acc, ok := e["accumulate_points"].(map[string]interface{}); ok {
			if p, ok := acc["points"].(float64); ok {
				points = int64(p)
			}
		} else if red, ok := e["redeem_points"].(map[string]interface{}); ok {
			if p, ok := red["points"].(float64); ok {
				points = int64(p)
			}
		}

		history = append(history, map[string]interface{}{
			"type":   eventType,
			"points": points,
			"date":   createdAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"transactions": history})
}
