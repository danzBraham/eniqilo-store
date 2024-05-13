package repositories

import "github.com/danzBraham/eniqilo-store/internal/domains/entities"

type StaffRepository interface {
	CreateStaff(*entities.RegisterStaff) (id string, err error)
	VerifyPhoneNumber(phoneNumber string) (bool, error)
	FindByPhoneNumber(phoneNumber string) (*entities.Staff, error)
}
