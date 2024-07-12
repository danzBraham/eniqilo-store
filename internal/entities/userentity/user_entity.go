package userentity

type Role string

const (
	Staff    Role = "Staff"
	Customer Role = "Customer"
)

type User struct {
	ID          string
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
	UserID      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type LoginStaffRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,e164"`
	Password    string `json:"password" validate:"required,min=5,max=15"`
}

type LoginStaffResponse struct {
	UserID      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type RegisterCustomerRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,e164"`
	Name        string `json:"name" validate:"required,min=5,max=50"`
}

type RegisterCustomerResponse struct {
	UserID      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}

type CustomerQueryParams struct {
	PhoneNumber string
	Name        string
}

type GetCustomerResponse struct {
	UserID      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}
