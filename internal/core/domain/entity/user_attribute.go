package entity

import "github.com/google/uuid"

type UserAttribute struct {
	UserID      uuid.UUID `json:"user_id"`
	AttributeID uuid.UUID `json:"attribute_id"`
	View        bool      `json:"view"`
	Search      bool      `json:"search"`
	Detail      bool      `json:"detail"`
	Add         bool      `json:"add"`
	Update      bool      `json:"update"`
	Delete      bool      `json:"delete"`
	Export      bool      `json:"export"`
	Import      bool      `json:"import"`
	CanSeePrice bool      `json:"can_see_price"`
}
