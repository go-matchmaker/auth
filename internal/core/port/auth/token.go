package auth

import (
	"auth/internal/core/domain/entity"
	"auth/internal/core/domain/valueobject"
	"time"
)

type TokenMaker interface {
	CreateToken(userID string, name, surname, email, role, phoneNumber, departmentID string, createdAt time.Time, attributes map[string]*entity.Permission) (string, string, *valueobject.Payload, error)
	CreateRefreshToken(payload *valueobject.Payload) (string, string, error)
	DecodeToken(token, public string) (*valueobject.Payload, error)
}
