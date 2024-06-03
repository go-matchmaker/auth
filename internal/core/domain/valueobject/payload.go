package valueobject

import (
	"auth/internal/core/domain/entity"
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
		ID           string                        `json:"id"`
		Name         string                        `json:"name"`
		Surname      string                        `json:"surname"`
		Email        string                        `json:"email"`
		Role         string                        `json:"role"`
		PhoneNumber  string                        `json:"phone_number"`
		DepartmentID string                        `json:"department_id"`
		CreatedAt    time.Time                     `json:"created_at"`
		IssuedAt     time.Time                     `json:"issued_at"`
		ExpiredAt    time.Time                     `json:"expired_at"`
		Attributes   map[string]*entity.Permission `json:"attributes"`
	}
)

func NewPayload(userID string, name, surname, email, role, phoneNumber, departmentID string, createdAt time.Time, attributes map[string]*entity.Permission, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		ID:           userID,
		Role:         role,
		Email:        email,
		Name:         name,
		Surname:      surname,
		PhoneNumber:  phoneNumber,
		DepartmentID: departmentID,
		Attributes:   attributes,
		CreatedAt:    createdAt,
		IssuedAt:     time.Now(),
		ExpiredAt:    time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if !time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
