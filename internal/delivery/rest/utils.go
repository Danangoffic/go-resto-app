package rest

import (
	"errors"
	"net/http"
	"resto-app/internal/model"
	"strings"
)

func GetSessionData(r *http.Request) (model.UserSession, error) {
	authString := r.Header.Get("Authorization")
	splitString := strings.Split(authString, " ")
	if len(splitString) != 2 {
		return model.UserSession{}, errors.New("unauthorized")
	}

	if splitString[0] != "Bearer" {
		return model.UserSession{}, errors.New("Auhtorization declined")
	}

	accessString := splitString[1]
	return model.UserSession{
		JWTToken: accessString,
	}, nil
}
