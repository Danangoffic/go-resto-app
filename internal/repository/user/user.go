package user

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"errors"
	"resto-app/internal/model"
	"resto-app/internal/tracing"
	"time"

	"gorm.io/gorm"
)

type userRepo struct {
	db        *gorm.DB
	gcm       cipher.AEAD
	time      uint32
	memory    uint32
	threads   uint8
	keyLen    uint32
	signKey   *rsa.PrivateKey
	accessExp time.Duration
}

func GetRepository(
	db *gorm.DB,
	secret string,
	time uint32,
	memory uint32,
	threads uint8,
	keyLen uint32,
	signKey *rsa.PrivateKey,
	accessExp time.Duration,
) (Repository, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &userRepo{
		db:        db,
		gcm:       gcm,
		time:      time,
		memory:    memory,
		threads:   threads,
		keyLen:    keyLen,
		signKey:   signKey,
		accessExp: accessExp,
	}, nil
}

func (r *userRepo) RegisterUser(ctx context.Context, userData model.User) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()

	if err := r.db.WithContext(ctx).Create(&userData).Error; err != nil {
		return model.User{}, err
	}

	return userData, nil
}

func (r *userRepo) CheckRegistered(ctx context.Context, username string) (bool, error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckRegistered")
	defer span.End()

	var userData model.User
	err := r.db.WithContext(ctx).Where(model.User{Username: username}).First(&userData).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return userData.ID != "", nil
}

func (r *userRepo) ValidateUser(ctx context.Context, username string, password string) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "ValidateUser")
	defer span.End()

	var userData model.User
	hashPassword, err := r.GenerateUserHash(ctx, password)
	if err != nil {
		return model.User{}, err
	}
	err = r.db.WithContext(ctx).Where(model.User{Username: username, Hash: hashPassword}).First(&userData).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.User{}, errors.New("User not found")
		}
		return model.User{}, err
	}
	return userData, nil
}

func (r *userRepo) GetUserData(ctx context.Context, username string) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetUserData")
	defer span.End()

	var userData model.User
	err := r.db.WithContext(ctx).Where(model.User{Username: username}).First(&userData).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return userData, nil
		}
		return userData, err
	}
	return userData, nil
}

func (r *userRepo) VerifyLogin(ctx context.Context, username, password string, userData model.User) (bool, error) {
	ctx, span := tracing.CreateSpan(ctx, "VerifyLogin")
	defer span.End()

	if username != userData.Username {
		return false, nil
	}
	verified, err := r.comparePassword(ctx, password, userData.Hash)
	if err != nil {
		return false, err
	}

	return verified, nil
}
