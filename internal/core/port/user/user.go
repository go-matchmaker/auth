package user

import (
	"auth/internal/core/domain/aggregate"
	"auth/internal/core/domain/entity"
	"context"
)

type UserRepositoryPort interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserPassword(ctx context.Context, email string) (string, error)
}

type UserServicePort interface {
	Login(ctx context.Context, email, password string) (*aggregate.UserAccess, error)
	GetUser(ctx context.Context, id string) (*entity.User, error)
}
