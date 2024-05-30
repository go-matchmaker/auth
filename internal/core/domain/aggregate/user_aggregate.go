package aggregate

import (
	"auth/internal/core/domain/entity"
)

type UserAggregate struct {
	User            entity.User             `json:"user"`
	Role            entity.Role             `json:"role"`
	Departments     []entity.Department     `json:"departments"`
	Attributes      []entity.Attribute      `json:"attributes"`
	UserAttributes  []entity.UserAttribute  `json:"user_attributes"`
	UserDepartments []entity.UserDepartment `json:"user_departments"`
}

func NewUserAggregate(user entity.User, role entity.Role, departments []entity.Department, attributes []entity.Attribute, userAttributes []entity.UserAttribute, userDepartments []entity.UserDepartment) *UserAggregate {
	return &UserAggregate{
		User:            user,
		Role:            role,
		Departments:     departments,
		Attributes:      attributes,
		UserAttributes:  userAttributes,
		UserDepartments: userDepartments,
	}
}
