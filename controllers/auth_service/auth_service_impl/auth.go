package authserviceimpl

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/MalikSaddique/chat_application_go/middleware"
	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var refreshSecretKey = []byte("my_refresh_secret_key")
var secretKey = []byte("secret-key")

func (u *AuthServiceImpl) SignUp(c *gin.Context, req *models.UserSignUp) *models.UserSignUp {

	createdUser := u.userAuth.SignUp(c, req)

	if createdUser == nil {
		fmt.Println("signup error")
		return nil
	}

	response := models.UserSignUp{
		// Id:       createdUser.Id,
		Email:    createdUser.Email,
		Password: createdUser.Password,
		Message:  "User created successfully",
	}

	return &response
}

func (u *AuthServiceImpl) Login(c *gin.Context, req *models.UserLogin) (*models.TokenPair, error) {

	user, err := u.userAuth.FindUserByEmail(req.Email)

	if err != nil {
		return nil, errors.New("User not found")
	}

	if user.Password != req.Password || user.Email != req.Email {
		return nil, errors.New("Invalid credentials")
	}

	token, err := middleware.CreateToken(user.Email, int(user.Id))
	if err != nil {
		return nil, errors.New("Failed to generate access token")
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString(refreshSecretKey)
	if err != nil {
		return nil, errors.New("Failed to generate refresh token")
	}

	response := models.TokenPair{
		AccessToken:  token,
		RefreshToken: refreshTokenString,
	}
	return &response, nil
}

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
