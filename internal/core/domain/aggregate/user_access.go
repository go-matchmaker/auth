package aggregate

import (
	"auth/internal/core/domain/entity"
	"auth/internal/core/domain/valueobject"
	"github.com/google/uuid"
	"time"
)

type (
	UserAcess struct {
		ID            uuid.UUID   `json:"id"`
		Role          entity.Role `json:"role"`
		Name          string      `json:"name"`
		Surname       string      `json:"surname"`
		Email         string      `json:"email"`
		PhoneNumber   string      `json:"phone_number"`
		PasswordHash  string      `json:"password_hash"`
		CreatedAt     time.Time   `json:"created_at"`
		AccessToken   string      `json:"access_token"`
		AccessPublic  string      `json:"access_public"`
		RefreshToken  string      `json:"refresh_token"`
		RefreshPublic string      `json:"refresh_public"`
		ClientIp      string      `json:"client_ip"`
		IsBlocked     bool        `json:"is_blocked"`
		IssuedAt      time.Time   `json:"issued_at"`
		ExpiredAt     time.Time   `json:"expires_at"`
	}
)

func NewUserAcess(user *entity.User, payload *valueobject.Payload, accessToken, accessPublic, refreshToken, refreshPublic, clientIp string) *UserAcess {
	return &UserAcess{
		ID:            user.ID,
		Role:          user.Role,
		Name:          user.Name,
		Surname:       user.Surname,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		PasswordHash:  user.PasswordHash,
		CreatedAt:     user.CreatedAt,
		AccessToken:   accessToken,
		AccessPublic:  accessPublic,
		RefreshToken:  refreshToken,
		RefreshPublic: refreshPublic,
		ClientIp:      clientIp,
		IsBlocked:     payload.IsBlocked,
		IssuedAt:      payload.IssuedAt,
		ExpiredAt:     payload.ExpiredAt,
	}
}
