package interfaces

import "github.com/danzBraham/eniqilo-store/internal/domains/entities"

type StaffService interface {
	RegisterStaff(*entities.RegisterStaff) (*entities.RegisteredStaff, error)
}
