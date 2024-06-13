package aggregate

import (
	"auth/internal/core/domain/entity"
)

type (
	UserAccess struct {
		ID            string `json:"id"`
		AccessToken   string `json:"access_token"`
		AccessPublic  string `json:"access_public"`
		RefreshToken  string `json:"refresh_token"`
		RefreshPublic string `json:"refresh_public"`
	}
)

func NewUserAccess(user *entity.User, accessToken, accessPublic, refreshToken, refreshPublic string) *UserAccess {
	return &UserAccess{
		ID:            user.ID,
		AccessToken:   accessToken,
		AccessPublic:  accessPublic,
		RefreshToken:  refreshToken,
		RefreshPublic: refreshPublic,
	}
}
