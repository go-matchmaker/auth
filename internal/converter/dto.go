package converter

import (
	"auth/internal/core/domain/entity"
	"auth/internal/dto"
	"time"
)

func UserRegisterRequestToModel(userDto *dto.UserRegisterRequest, role, pass string) (*entity.User, error) {
	return &entity.User{
		Role:        entity.Role(role),
		Name:        userDto.Name,
		Surname:     userDto.Surname,
		Email:       userDto.Email,
		PhoneNumber: userDto.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
