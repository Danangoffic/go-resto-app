package user

import (
	"errors"
	"resto-app/internal/model"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
}

func (r *userRepo) CreateUserSession(userID string) (model.UserSession, error) {
	accessToken, err := r.generateAccessToken(userID)

	if err != nil {
		return model.UserSession{}, err
	}

	return model.UserSession{
		JWTToken: accessToken,
	}, nil
}

func (r *userRepo) generateAccessToken(userID string) (string, error) {
	accessTokenExp := time.Now().Add(r.accessExp).Unix()
	accessClaims := Claims{
		jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: accessTokenExp,
		},
	}

	accessJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), accessClaims)

	return accessJwt.SignedString(r.signKey)
}

func (r *userRepo) CheckSession(data model.UserSession) (userID string, err error) {
	accessToken, err := jwt.ParseWithClaims(data.JWTToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return &r.signKey.PublicKey, nil
	})
	if err != nil {
		return "", nil
	}

	accessTokenClaim, ok := accessToken.Claims.(*Claims)
	if !ok {
		return "", errors.New("unauthorized")
	}

	if accessToken.Valid {
		return accessTokenClaim.Subject, nil
	}

	return "", errors.New("unauthorized")
}
