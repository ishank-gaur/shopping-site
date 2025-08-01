package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "ecommerce/utils" // adjust import path
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
            return
        }

        token := parts[1]
        userID, err := utils.ValidateJWT(token)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired access token"})
            return
        }

        // Attach user ID to context
        c.Set("userID", userID)
        c.Next()
    }
}
