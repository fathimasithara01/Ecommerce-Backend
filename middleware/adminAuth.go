package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/fathimasithara01/ecommerce/config"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config, _ := config.LoadConfig()
		jwtSecret := []byte(config.Key)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or incorrect"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"meaage": "invalid authorization format"})
			return
		}

		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid claims"})
			return
		}
		// Check admin role
		// role, ok := claims["role"].(string)
		// if !ok || role != "admin" {
		// 	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Admin access required"})
		// 	return
		// }

		// Optional: store admin ID in context
		c.Set("admin_id", claims["user_id"])
		c.Next()
	}
}
