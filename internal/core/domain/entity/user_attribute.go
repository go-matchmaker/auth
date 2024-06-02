package entity

type UserAttribute struct {
	UserID      string `json:"user_id"`
	Attribute   string `json:"attribute_id"`
	View        bool   `json:"view"`
	Search      bool   `json:"search"`
	Detail      bool   `json:"detail"`
	Add         bool   `json:"add"`
	Update      bool   `json:"update"`
	Delete      bool   `json:"delete"`
	Export      bool   `json:"export"`
	Import      bool   `json:"import"`
	CanSeePrice bool   `json:"can_see_price"`
}
