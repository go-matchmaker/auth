package valueobject

import (
	"errors"
	"time"
)

const (
	AccessToken  = "access"
	RefreshToken = "refresh"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type (
	//Attributes going to add
	Payload struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Surname   string    `json:"surname"`
		Email     string    `json:"email"`
		Role      string    `json:"role"`
		IsBlocked bool      `json:"is_blocked"`
		IssuedAt  time.Time `json:"issued_at"`
		ExpiredAt time.Time `json:"expired_at"`
	}
)

func NewPayload(userID string, email, role string, isBlocked bool, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		ID:        userID,
		Role:      role,
		Email:     email,
		IsBlocked: isBlocked,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if !time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
