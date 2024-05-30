package dto

import (
	"auth/internal/core/domain/aggregate"
	"time"
)

type UserLoginRequestResponse struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Surname       string    `json:"surname"`
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phone_number"`
	Role          string    `json:"role"`
	CreatedAt     time.Time `json:"created_at"`
	AccessToken   string    `json:"access_token"`
	AccessPublic  string    `json:"access_public"`
	RefreshToken  string    `json:"refresh_token"`
	RefreshPublic string    `json:"refresh_public"`
	ExpiredAt     time.Time `json:"expired_at"`
}

func NewUserLoginRequestResponse(userData *aggregate.UserAcess) *UserLoginRequestResponse {
	return &UserLoginRequestResponse{
		ID:            userData.ID.String(),
		Name:          userData.Name,
		Surname:       userData.Surname,
		Email:         userData.Email,
		PhoneNumber:   userData.PhoneNumber,
		Role:          string(userData.Role),
		CreatedAt:     userData.CreatedAt,
		AccessToken:   userData.AccessToken,
		AccessPublic:  userData.AccessPublic,
		RefreshToken:  userData.RefreshToken,
		RefreshPublic: userData.RefreshPublic,
		ExpiredAt:     userData.ExpiredAt,
	}
}
