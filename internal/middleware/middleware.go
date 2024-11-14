package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"mygo/internal/auth"
)

// /////////////////////////////////////////////////////////////////////////
func JWTMiddleware(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	claims, err := auth.ValidateJWT(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
		c.Abort()
		return
	}

	c.Set("username", claims.Username)

	c.Next()
}
