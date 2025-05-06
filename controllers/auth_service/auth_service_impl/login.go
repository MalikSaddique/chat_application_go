package authserviceimpl

import (
	"errors"
	"time"

	"github.com/MalikSaddique/chat_application_go/middleware"
	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

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
