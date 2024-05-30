package dto

type UserRegisterRequest struct {
	Email        string                 `json:"email" binding:"required,email"`
	Password     string                 `json:"password" binding:"required,min=8"`
	Name         string                 `json:"name" binding:"required"`
	Surname      string                 `json:"surname" binding:"required"`
	PhoneNumber  string                 `json:"phone_number"`
	DepartmentID int                    `json:"department_id"`
	Attributes   map[string]Permissions `json:"attributes"`
}
type Permissions struct {
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

type User struct {
	Role        string `json:"role"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type PasswordChangeRequest struct {
	OldPassword string `json:"old_password" binding:"required,min=8"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}
