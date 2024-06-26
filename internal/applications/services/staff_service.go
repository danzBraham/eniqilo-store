package services

import (
	"fmt"

	"github.com/danzBraham/eniqilo-store/internal/applications/interfaces"
	"github.com/danzBraham/eniqilo-store/internal/domains/entities"
	"github.com/danzBraham/eniqilo-store/internal/domains/repositories"
	"github.com/danzBraham/eniqilo-store/internal/helpers"
)

type StaffService struct {
	StaffRepository repositories.StaffRepository
}

func NewStaffService(staffRepository repositories.StaffRepository) interfaces.StaffService {
	return &StaffService{StaffRepository: staffRepository}
}

func (s *StaffService) RegisterStaff(staff *entities.RegisterStaff) (*entities.LoggedInStaff, error) {
	// Check if phone number already exists
	if isPhoneNumberExists, _ := s.StaffRepository.VerifyPhoneNumber(staff.PhoneNumber); isPhoneNumberExists {
		return nil, fmt.Errorf("staff with phone number %s already exists", staff.PhoneNumber)
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(staff.Password)
	if err != nil {
		return nil, err
	}
	staff.Password = hashedPassword

	// If it doesn't create the new staff
	id, err := s.StaffRepository.CreateStaff(staff)
	if err != nil {
		return nil, err
	}

	// Create access token
	accessToken, err := helpers.CreateJWT(id)
	if err != nil {
		return nil, err
	}

	return &entities.LoggedInStaff{
		ID:          id,
		PhoneNumber: staff.PhoneNumber,
		Name:        staff.Name,
		AccessToken: accessToken,
	}, nil
}

func (s *StaffService) LoginStaff(staff *entities.LoginStaff) (*entities.LoggedInStaff, error) {
	// Find staff with phone number
	foundedStaff, err := s.StaffRepository.FindByPhoneNumber(staff.PhoneNumber)
	if err != nil {
		return nil, err
	}

	// Create access token
	accessToken, err := helpers.CreateJWT(foundedStaff.ID)
	if err != nil {
		return nil, err
	}

	return &entities.LoggedInStaff{
		ID:          foundedStaff.ID,
		PhoneNumber: staff.PhoneNumber,
		Name:        foundedStaff.Name,
		AccessToken: accessToken,
	}, nil
}
