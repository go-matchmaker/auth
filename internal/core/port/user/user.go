package user

import (
	"auth/internal/core/domain/aggregate"
	"auth/internal/core/domain/entity"
	"context"

	"github.com/google/uuid"
)

type UserRepositoryPort interface {
	Insert(ctx context.Context, user *entity.User) (*uuid.UUID, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdatePassword(ctx context.Context, id uuid.UUID, password string) (*entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	DeleteOne(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context) error
}

type UserServicePort interface {
	Register(ctx context.Context, req *entity.User) (*uuid.UUID, error)
	Login(ctx context.Context, email, password, ip string) (*aggregate.UserAcess, error)
	UserSelfInfo(ctx context.Context, id string) (entity.User, error)
	GetUserWithPermissions(ctx context.Context, id string) (*entity.UserWithPermissions, error)
	//DeleteOne(ctx context.Context, id uuid.UUID) error
	// GetByID(ctx context.Context, id string) (*model.User, error)
	// RefreshToken(ctx context.Context, userID string) (string, error)
	// ChangePassword(ctx context.Context, id string, req *dto.ChangePasswordReq) error
}
