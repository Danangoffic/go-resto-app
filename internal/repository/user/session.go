package user

import (
	"context"
	"errors"
	"resto-app/internal/model"
	"resto-app/internal/tracing"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
}

func (r *userRepo) CreateUserSession(ctx context.Context, userID string) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateUserSession")
	defer span.End()

	accessToken, err := r.generateAccessToken(ctx, userID)

	if err != nil {
		return model.UserSession{}, err
	}

	return model.UserSession{
		JWTToken: accessToken,
	}, nil
}

func (r *userRepo) generateAccessToken(ctx context.Context, userID string) (string, error) {
	ctx, span := tracing.CreateSpan(ctx, "generateAccessToken")
	defer span.End()

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

func (r *userRepo) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckSession")
	defer span.End()

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
