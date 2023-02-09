package user

import "resto-app/internal/model"

type Repository interface {
	RegisterUser(userData model.User) (model.User, error)
	CheckRegistered(username string) (bool, error)
	GenerateUserHash(password string) (hash string, err error)
	ValidateUser(username string, password string) (model.User, error)
	GetUserData(username string) (model.User, error)
	CreateUserSession(userID string) (model.UserSession, error)
	VerifyLogin(username, password string, userData model.User) (bool, error)
	CheckSession(data model.UserSession) (userID string, err error)
}
