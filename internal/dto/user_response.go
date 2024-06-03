package dto

import (
	"time"
)

type UserLoginRequestResponse struct {
	Token     string      `json:"token"`
	PublicKey string      `json:"public_key"`
	User      *UserAccess `json:"user"`
}

type UserAccess struct {
	ID            string `json:"id"`
	AccessToken   string `json:"access_token"`
	AccessPublic  string `json:"access_public"`
	RefreshToken  string `json:"refresh_token"`
	RefreshPublic string `json:"refresh_public"`
}

type GetUserResponse struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Surname      string                 `json:"surname"`
	Email        string                 `json:"email"`
	PhoneNumber  string                 `json:"phone_number"`
	Role         string                 `json:"role"`
	DepartmentID string                 `json:"department_id"`
	Attributes   map[string]*Permission `json:"attributes"`
	CreatedAt    time.Time              `json:"created_at"`
}

type Permission struct {
	View        bool `json:"view"`
	Search      bool `json:"search"`
	Detail      bool `json:"detail"`
	Add         bool `json:"add"`
	Update      bool `json:"update"`
	Delete      bool `json:"delete"`
	Export      bool `json:"export"`
	Import      bool `json:"import"`
	CanSeePrice bool `json:"can_see_price"`
}
