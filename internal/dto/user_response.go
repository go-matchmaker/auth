package dto

import (
	"auth/internal/core/domain/aggregate"
	"auth/internal/core/domain/entity"
	"time"
)

type UserLoginRequestResponse struct {
	Token     string                `json:"token"`
	PublicKey string                `json:"public_key"`
	User      *aggregate.UserAccess `json:"user"`
}

type GetUserRequestResponse struct {
	ID          string                       `json:"id"`
	Name        string                       `json:"name"`
	Surname     string                       `json:"surname"`
	Email       string                       `json:"email"`
	PhoneNumber string                       `json:"phone_number"`
	Role        string                       `json:"role"`
	CreatedAt   time.Time                    `json:"created_at"`
	Attributes  map[string]entity.Permission `json:"attributes"`
}

func NewUserLoginRequestResponse(userData *aggregate.UserAccess) *UserLoginRequestResponse {
	return &UserLoginRequestResponse{
		Token:     userData.AccessToken,
		PublicKey: userData.AccessPublic,
		User:      userData,
	}
}

func NewUserResponse(userData *entity.User) *GetUserRequestResponse {
	return &GetUserRequestResponse{
		ID:          userData.ID,
		Name:        userData.Name,
		Surname:     userData.Surname,
		Email:       userData.Email,
		PhoneNumber: userData.PhoneNumber,
		Role:        string(userData.Role),
		CreatedAt:   userData.CreatedAt,
		Attributes:  userData.UserPermissions,
	}
}
