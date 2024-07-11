package userentity

type Role string

const (
	Staff    Role = "Staff"
	Customer Role = "Customer"
)

type User struct {
	Id          string
	PhoneNumber string
	Name        string
	Password    string
	Role        Role
	CreatedAt   string
	UpdatedAt   string
}

type RegisterStaffRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,e164"`
	Name        string `json:"name" validate:"required,min=5,max=50"`
	Password    string `json:"password" validate:"required,min=5,max=15"`
}

type RegisterStaffResponse struct {
	UserId      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
