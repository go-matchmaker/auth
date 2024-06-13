package converter

import (
	"auth/internal/core/domain/aggregate"
	"auth/internal/core/domain/entity"
	"auth/internal/dto"
)

func UserLoginModelToDto(userData *aggregate.UserAccess) *dto.UserLoginRequestResponse {
	return &dto.UserLoginRequestResponse{
		Token:     userData.AccessToken,
		PublicKey: userData.AccessPublic,
		User:      (*dto.UserAccess)(userData),
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
		Attributes:   userPermissionsModelToDto(userData.Permissions),
	}
}

func userPermissionsModelToDto(permissions map[string]*entity.Permission) map[string]*dto.Permission {
	result := make(map[string]*dto.Permission)
	for key, value := range permissions {
		result[key] = &dto.Permission{
			View:        value.View,
			Search:      value.Search,
			Detail:      value.Detail,
			Add:         value.Add,
			Update:      value.Update,
			Delete:      value.Delete,
			Export:      value.Export,
			Import:      value.Import,
			CanSeePrice: value.CanSeePrice,
		}
	}
	return result
}
