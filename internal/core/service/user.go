package service

import (
	"auth/internal/core/domain/aggregate"
	"auth/internal/core/domain/entity"
	"auth/internal/core/port/auth"
	"auth/internal/core/port/user"
	"auth/internal/core/util"
	"context"
	"errors"
	"github.com/google/uuid"
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

func (us *UserService) Register(ctx context.Context, userModel *entity.User) (*uuid.UUID, error) {
	id, err := us.userRepo.Insert(ctx, userModel)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (us *UserService) Login(ctx context.Context, email, password, ip string) (*aggregate.UserAcess, error) {
	userModel, err := us.userRepo.GetByEmail(ctx, email)
	sessionModel := new(aggregate.UserAcess)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = util.ComparePassword(password, userModel.PasswordHash)
	if err != nil {
		return nil, errors.New("password not match")
	}

	// We will take it from db later.
	isBlocked := false
	accessToken, publicKey, accessPayload, err := us.token.CreateToken(userModel.ID, userModel.Email, string(userModel.Role), isBlocked)
	if err != nil {
		return nil, err
	}
	refreshToken, refreshPublicKey, refreshPayload, err := us.token.CreateRefreshToken(accessPayload)
	if err != nil {
		return nil, err
	}

	sessionModel = aggregate.NewUserAcess(&userModel, refreshPayload, accessToken, publicKey, refreshToken, refreshPublicKey, ip)

	return sessionModel, nil
}

func (us *UserService) UserSelfInfo(ctx context.Context, email string) (entity.User, error) {
	userUpdate, err := us.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return entity.User{}, err
	}

	return userUpdate, nil
}
