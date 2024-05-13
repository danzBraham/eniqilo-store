package entities

import "time"

type Staff struct {
	ID          string    `json:"id"`
	PhoneNumber string    `json:"phoneNumber" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RegisterStaff struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16"`
	Name        string `json:"name" validate:"required,min=5,max=50"`
	Password    string `json:"password" validate:"required,min=5,max=15"`
}

type LoginStaff struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16"`
	Password    string `json:"password" validate:"required,min=5,max=15"`
}

type LoggedInStaff struct {
	ID          string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
