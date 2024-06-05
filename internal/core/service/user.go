package service

import (
	"auth/internal/core/domain/aggregate"
	"auth/internal/core/port/auth"
	"auth/internal/core/port/user"
	"auth/internal/core/util"
	"context"
	"errors"
	"fmt"

	"github.com/google/wire"
)

var (
	_              user.UserServicePort = (*UserService)(nil)
	UserServiceSet                      = wire.NewSet(NewUserService)
)

type UserService struct {
	userRepo user.UserRepositoryPort
	token    auth.TokenMaker
}

func NewUserService(userRepo user.UserRepositoryPort, token auth.TokenMaker) user.UserServicePort {
	return &UserService{
		userRepo,
		token,
	}
}

func (us *UserService) Login(ctx context.Context, email, password string) (*aggregate.UserAccess, error) {
	userPassword, err := us.userRepo.GetUserPassword(ctx, email)
	if err != nil {
		return nil, err
	}
	fmt.Println("userPassword", userPassword)
	err = util.ComparePassword(password, userPassword)
	if err != nil {
		return nil, errors.New("password not match")
	}
	fmt.Println("password match")
	userModel, err := us.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	accessToken, publicKey, accessPayload, err := us.token.CreateToken(userModel.User.ID)
	if err != nil {
		return nil, err
	}
	refreshToken, refreshPublicKey, err := us.token.CreateRefreshToken(accessPayload)
	if err != nil {
		return nil, err
	}

	sessionModel := aggregate.NewUserAccess(userModel.User, accessToken, publicKey, refreshToken, refreshPublicKey)

	return sessionModel, nil
}

func (us *UserService) GetUser(ctx context.Context, id string) (*aggregate.UserAggregate, error) {
	userModel, err := us.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return userModel, nil
}
