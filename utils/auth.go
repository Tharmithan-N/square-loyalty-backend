package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	fakeUserToken = "fake-jwt-token"
)

// Middleware to protect routes (checks Authorization header)
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "Bearer "+fakeUserToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Fake login handler
func FakeLogin(c *gin.Context) {
	// Accept any username/password in body for demo
	type loginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input loginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// You can add logic to validate username/password here.
	// For now, always succeed.

	c.JSON(http.StatusOK, gin.H{
		"token": fakeUserToken,
	})
}
