package authserviceimpl

import (
	"github.com/MalikSaddique/chat_application_go/database"
	authservice "github.com/MalikSaddique/chat_application_go/service/auth_service"
)

type AuthServiceImpl struct {
	userAuth database.Storage
}

func NewAuthService(input NewAuthServiceImpl) authservice.AuthService {
	return &AuthServiceImpl{
		userAuth: input.UserAuth,
	}
}

type NewAuthServiceImpl struct {
	UserAuth database.Storage
}
