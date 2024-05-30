package user

import (
	"auth/internal/core/domain/entity"
	"context"

	"github.com/google/uuid"
)

type AttributeRepositoryPort interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (entity.User, error)
}

type AttributeServicePort interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id string) (entity.User, error)
}
