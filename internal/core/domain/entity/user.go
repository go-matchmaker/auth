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
