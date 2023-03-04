package user

import (
	"context"
	"resto-app/internal/model"
)

type Repository interface {
	RegisterUser(ctx context.Context, userData model.User) (model.User, error)
	CheckRegistered(ctx context.Context, username string) (bool, error)
	GenerateUserHash(ctx context.Context, password string) (hash string, err error)
	ValidateUser(ctx context.Context, username string, password string) (model.User, error)
	GetUserData(ctx context.Context, username string) (model.User, error)
	CreateUserSession(ctx context.Context, userID string) (model.UserSession, error)
	VerifyLogin(ctx context.Context, username, password string, userData model.User) (bool, error)
	CheckSession(ctx context.Context, data model.UserSession) (userID string, err error)
}
