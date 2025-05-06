package authserviceimpl

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var refreshSecretKey = []byte("my_refresh_secret_key")
var secretKey = []byte("secret-key")

func (a *AuthServiceImpl) RefreshAccessToken(c *gin.Context) (string, error) {
	refreshTokenString := c.GetHeader("Authorization")
	refreshTokenString = strings.TrimPrefix(refreshTokenString, "Bearer ")

	if refreshTokenString == "" {
		return "", fmt.Errorf("Refresh token is empty")
	}

	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshSecretKey), nil
	})
	if err != nil || !refreshToken.Valid {
		return "", fmt.Errorf("Invalid refresh token")
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Invalid refresh token claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid email in refresh token")
	}

	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	})

	newAccessTokenString, err := newAccessToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("Failed to generate new access token")
	}

	return newAccessTokenString, nil
}
