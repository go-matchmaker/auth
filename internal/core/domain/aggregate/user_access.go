package aggregate

import (
	"auth/internal/core/domain/entity"
	"auth/internal/core/domain/valueobject"
	"time"
)

type (
	UserAccess struct {
		ID            string                       `json:"id"`
		Role          entity.Role                  `json:"role"`
		Name          string                       `json:"name"`
		Surname       string                       `json:"surname"`
		Email         string                       `json:"email"`
		PhoneNumber   string                       `json:"phone_number"`
		CreatedAt     time.Time                    `json:"created_at"`
		Attributes    map[string]entity.Permission `json:"attributes"`
		AccessToken   string                       `json:"access_token"`
		AccessPublic  string                       `json:"access_public"`
		RefreshToken  string                       `json:"refresh_token"`
		RefreshPublic string                       `json:"refresh_public"`
		IssuedAt      time.Time                    `json:"issued_at"`
		ExpiredAt     time.Time                    `json:"expires_at"`
	}
)

func NewUserAccess(user *entity.User, payload *valueobject.Payload, accessToken, accessPublic, refreshToken, refreshPublic string) *UserAccess {
	return &UserAccess{
		ID:            user.ID,
		Role:          user.Role,
		Name:          user.Name,
		Surname:       user.Surname,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		CreatedAt:     user.CreatedAt,
		AccessToken:   accessToken,
		AccessPublic:  accessPublic,
		RefreshToken:  refreshToken,
		RefreshPublic: refreshPublic,
		IssuedAt:      payload.IssuedAt,
		ExpiredAt:     payload.ExpiredAt,
		Attributes:    user.UserPermissions,
	}
}
