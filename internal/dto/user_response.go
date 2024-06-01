package dto

import (
	"auth/internal/core/domain/aggregate"
	"auth/internal/core/domain/entity"
	"time"
)

type UserLoginRequestResponse struct {
	ID            string                       `json:"id"`
	Name          string                       `json:"name"`
	Surname       string                       `json:"surname"`
	Email         string                       `json:"email"`
	PhoneNumber   string                       `json:"phone_number"`
	Role          string                       `json:"role"`
	CreatedAt     time.Time                    `json:"created_at"`
	AccessToken   string                       `json:"access_token"`
	AccessPublic  string                       `json:"access_public"`
	RefreshToken  string                       `json:"refresh_token"`
	RefreshPublic string                       `json:"refresh_public"`
	ExpiredAt     time.Time                    `json:"expired_at"`
	Attributes    map[string]entity.Permission `json:"attributes"`
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
		ID:            userData.ID,
		Name:          userData.Name,
		Surname:       userData.Surname,
		Email:         userData.Email,
		PhoneNumber:   userData.PhoneNumber,
		Role:          string(userData.Role),
		Attributes:    userData.Attributes,
		CreatedAt:     userData.CreatedAt,
		AccessToken:   userData.AccessToken,
		AccessPublic:  userData.AccessPublic,
		RefreshToken:  userData.RefreshToken,
		RefreshPublic: userData.RefreshPublic,
		ExpiredAt:     userData.ExpiredAt,
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
	}
}
