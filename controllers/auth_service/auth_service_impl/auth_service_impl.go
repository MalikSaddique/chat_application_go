package authserviceimpl

import (
	authservice "github.com/MalikSaddique/chat_application_go/controllers/auth_service"
	"github.com/MalikSaddique/chat_application_go/db"
)

type AuthServiceImpl struct {
	userAuth db.Storage
}

func NewAuthService(input NewAuthServiceImpl) authservice.AuthService {
	return &AuthServiceImpl{
		userAuth: input.UserAuth,
	}
}

type NewAuthServiceImpl struct {
	UserAuth db.Storage
}

var _ authservice.AuthService = &AuthServiceImpl{}
