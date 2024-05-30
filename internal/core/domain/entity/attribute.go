package entity

import (
	"github.com/google/uuid"
	"time"
)

type Attribute struct {
	ID           uuid.UUID `json:"id"`
	DepartmentID int       `json:"department_id"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
}
