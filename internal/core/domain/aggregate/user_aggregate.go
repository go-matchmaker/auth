package aggregate

import "auth/internal/core/domain/entity"

type UserAggregate struct {
	User        *entity.User                  `json:"user"`
	Department  *entity.Department            `json:"department"`
	Permissions map[string]*entity.Permission `json:"user_permission"`
}

func NewUserAggregate(user *entity.User, department *entity.Department, Permissions map[string]*entity.Permission) *UserAggregate {
	return &UserAggregate{
		User:        user,
		Department:  department,
		Permissions: Permissions,
	}
}
