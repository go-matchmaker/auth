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

type GetUserResponse struct {
	ID           string                        `json:"id"`
	Name         string                        `json:"name"`
	Surname      string                        `json:"surname"`
	Email        string                        `json:"email"`
	PhoneNumber  string                        `json:"phone_number"`
	Role         string                        `json:"role"`
	DepartmentID string                        `json:"department_id"`
	Attributes   map[string]*entity.Permission `json:"attributes"`
	CreatedAt    time.Time                     `json:"created_at"`
}
