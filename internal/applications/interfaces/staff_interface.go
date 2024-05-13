package interfaces

import "github.com/danzBraham/eniqilo-store/internal/domains/entities"

type StaffService interface {
	RegisterStaff(*entities.RegisterStaff) (*entities.LoggedInStaff, error)
	LoginStaff(*entities.LoginStaff) (*entities.LoggedInStaff, error)
}
