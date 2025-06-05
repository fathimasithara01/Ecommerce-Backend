package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/fathimasithara01/ecommerce/config"
)

func GenerateAdminJWT(userID uint, isAdmin bool) (string, error) {
	config, _ := config.LoadConfig()

	claims := jwt.MapClaims{
		"user_id":  userID,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err := token.SignedString([]byte(config.Key))
	// if err != nil {
	// 	return "", err
	// }
	jwtSecret := []byte(config.Key)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenStr string) (*jwt.MapClaims, error) {
	config, _ := config.LoadConfig()
	jwtSecret := []byte(config.Key)

	token, err := jwt.ParseWithClaims(tokenStr, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
