package user

import (
	"auth/internal/core/domain/aggregate"
	"context"
)

type UserRepositoryPort interface {
	GetByID(ctx context.Context, id string) (*aggregate.UserAggregate, error)
	GetByEmail(ctx context.Context, email string) (*aggregate.UserAggregate, error)
	GetUserPassword(ctx context.Context, email string) (string, error)
}

type UserServicePort interface {
	Login(ctx context.Context, email, password string) (*aggregate.UserAccess, error)
	GetUser(ctx context.Context, id string) (*aggregate.UserAggregate, error)
}
