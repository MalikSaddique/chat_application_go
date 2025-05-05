package authserviceimpl

import (
	authservice "github.com/MalikSaddique/chat_application_go/controllers/auth_service"
	db "github.com/MalikSaddique/chat_application_go/db/postgresDB"
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
