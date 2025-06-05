package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fathimasithara01/ecommerce/utils/helper"
	"github.com/fathimasithara01/ecommerce/utils/response"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := helper.GetTokenFromHeader(authHeader)

		if tokenString == "" {
			var err error
			tokenString, err = c.Cookie("Authorization")
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		userID, userEmail, err := helper.ExtractUserIDFromToken(tokenString)
		if err != nil {
			fmt.Println("error is ", err)
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token ", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		if userID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
			return
		}
		fmt.Println("userid", userID)
		c.Set("user_id", uint(userID))
		c.Set("user_email", userEmail)
		c.Next()
	}
}

// var jwtSecret = []byte("your_jwt_secret_key") // replace with your secret

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header missing"})
// 			return
// 		}

// 		parts := strings.SplitN(authHeader, " ", 2)
// 		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization header format"})
// 			return
// 		}

// 		tokenString := parts[1]

// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			// Validate signing method if needed
// 			return jwtSecret, nil
// 		})

// 		if err != nil || !token.Valid {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
// 			return
// 		}

// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
// 			return
// 		}

// 		// Extract user_id from claims (assuming it's stored as float64)
// 		uidFloat, ok := claims["user_id"].(float64)
// 		if !ok {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "User ID not found in token"})
// 			return
// 		}

// 		userID := uint(uidFloat)

// 		if userID == 0 {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid user ID in token"})
// 			return
// 		}

// 		// Set user_id in context for handlers
// 		c.Set("user_id", userID)

// 		c.Next()
// 	}
// }
