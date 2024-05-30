package auth

import (
	"auth/internal/core/domain/valueobject"
)

type TokenMaker interface {
	CreateToken(id string, email, role string, isBlocked bool) (string, string, *valueobject.Payload, error)
	CreateRefreshToken(payload *valueobject.Payload) (string, string, *valueobject.Payload, error)
	DecodeToken(token, public string) (*valueobject.Payload, error)
}
