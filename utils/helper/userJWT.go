package helper

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/fathimasithara01/ecommerce/config"
)

type AuthUserClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims

	// Role  string `json:"role"`
}

func GenerateJWT(userID uint, email string, role string) (string, error) {
	config, _ := config.LoadConfig()

	Claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)

	tokenString, err := token.SignedString([]byte(config.Key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetTokenFromHeader(header string) string {
	if len(header) >= 7 && header[:7] == "Bearer " {
		return header[7:]
	}
	return header
}

func ExtractUserIDFromToken(tokenString string) (uint, string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return 0, "", fmt.Errorf("config error: %w", err)
	}
	token, err := jwt.ParseWithClaims(tokenString, &AuthUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Key), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*AuthUserClaims)
	if !ok || !token.Valid {
		return 0, "", errors.New("invalid token claims")
	}

	return claims.UserID, claims.Email, nil
}
