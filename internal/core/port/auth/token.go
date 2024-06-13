package auth

import (
	"auth/internal/core/domain/valueobject"
)

type TokenMaker interface {
	CreateToken(userID string) (string, string, *valueobject.Payload, error)
	CreateRefreshToken(payload *valueobject.Payload) (string, string, error)
	DecodeToken(token, public string) (*valueobject.Payload, error)
}
