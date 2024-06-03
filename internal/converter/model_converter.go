package converter

import (
	"auth/internal/core/domain/aggregate"
	"auth/internal/dto"
)

func UserLoginModelToDto(userData *aggregate.UserAccess) *dto.UserLoginRequestResponse {
	return &dto.UserLoginRequestResponse{
		Token:     userData.AccessToken,
		PublicKey: userData.AccessPublic,
		User:      userData,
	}
}

func GetUserModelToDto(userData *aggregate.UserAggregate) *dto.GetUserResponse {
	return &dto.GetUserResponse{
		ID:           userData.User.ID,
		Name:         userData.User.Name,
		Surname:      userData.User.Surname,
		Email:        userData.User.Email,
		PhoneNumber:  userData.User.PhoneNumber,
		Role:         string(userData.User.Role),
		CreatedAt:    userData.User.CreatedAt,
		DepartmentID: userData.Department.ID,
		Attributes:   userData.Permissions,
	}
}
