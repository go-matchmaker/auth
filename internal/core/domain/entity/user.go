package entity

import (
	"time"
)

type (
	Role string
)

var (
	SuperAdminRole Role = "super_admin"
	AdminRole      Role = "admin"
	EmployeeRole   Role = "employee"
	UserRole       Role = "user"
)

type User struct {
	ID          string    `json:"id"`
	Role        Role      `json:"role"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Permission struct {
	View        bool `json:"view"`
	Search      bool `json:"search"`
	Detail      bool `json:"detail"`
	Add         bool `json:"add"`
	Update      bool `json:"update"`
	Delete      bool `json:"delete"`
	Export      bool `json:"export"`
	Import      bool `json:"import"`
	CanSeePrice bool `json:"can_see_price"`
}
