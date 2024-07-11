package services

import (
	"context"
	"time"

	"github.com/danzBraham/eniqilo-store/internal/entities/userentity"
	"github.com/danzBraham/eniqilo-store/internal/errors/usererror"
	"github.com/danzBraham/eniqilo-store/internal/helpers/bcrypt"
	"github.com/danzBraham/eniqilo-store/internal/helpers/jwt"
	"github.com/danzBraham/eniqilo-store/internal/repositories"
	"github.com/oklog/ulid/v2"
)

type UserService interface {
	CreateStaff(ctx context.Context, payload *userentity.RegisterStaffRequest) (*userentity.RegisterStaffResponse, error)
	LoginStaff(ctx context.Context, payload *userentity.LoginStaffRequest) (*userentity.LoginStaffResponse, error)
	RegisterCustomer(ctx context.Context, payload *userentity.RegisterCustomerRequest) (*userentity.RegisterCustomerResponse, error)
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (s *UserServiceImpl) CreateStaff(ctx context.Context, payload *userentity.RegisterStaffRequest) (*userentity.RegisterStaffResponse, error) {
	isPhoneNumberExists, err := s.UserRepository.IsPhoneNumberByItsRoleExists(ctx, payload.PhoneNumber, userentity.Staff)
	if err != nil {
		return nil, err
	}
	if isPhoneNumberExists {
		return nil, usererror.ErrPhoneNumberAlreadyExists
	}

	hashedPassword, err := bcrypt.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	staff := &userentity.User{
		ID:          ulid.Make().String(),
		PhoneNumber: payload.PhoneNumber,
		Name:        payload.Name,
		Password:    hashedPassword,
		Role:        userentity.Staff,
	}

	err = s.UserRepository.CreateUser(ctx, staff)
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(2*time.Hour, staff.ID)
	if err != nil {
		return nil, err
	}

	return &userentity.RegisterStaffResponse{
		UserID:      staff.ID,
		PhoneNumber: staff.PhoneNumber,
		Name:        staff.Name,
		AccessToken: token,
	}, nil
}

func (s *UserServiceImpl) LoginStaff(ctx context.Context, payload *userentity.LoginStaffRequest) (*userentity.LoginStaffResponse, error) {
	staff, err := s.UserRepository.GetUserByPhoneNumberAndRole(ctx, payload.PhoneNumber, userentity.Staff)
	if err != nil {
		return nil, err
	}

	err = bcrypt.VerifyPassword(staff.Password, payload.Password)
	if err != nil {
		return nil, usererror.ErrInvalidPassword
	}

	token, err := jwt.GenerateToken(2*time.Hour, staff.ID)
	if err != nil {
		return nil, err
	}

	return &userentity.LoginStaffResponse{
		UserID:      staff.ID,
		PhoneNumber: staff.PhoneNumber,
		Name:        staff.Name,
		AccessToken: token,
	}, nil
}

func (s *UserServiceImpl) RegisterCustomer(ctx context.Context, payload *userentity.RegisterCustomerRequest) (*userentity.RegisterCustomerResponse, error) {
	isPhoneNumberExists, err := s.UserRepository.IsPhoneNumberByItsRoleExists(ctx, payload.PhoneNumber, userentity.Customer)
	if err != nil {
		return nil, err
	}
	if isPhoneNumberExists {
		return nil, usererror.ErrPhoneNumberAlreadyExists
	}

	customer := &userentity.User{
		ID:          ulid.Make().String(),
		PhoneNumber: payload.PhoneNumber,
		Name:        payload.Name,
		Role:        userentity.Customer,
	}

	err = s.UserRepository.CreateUser(ctx, customer)
	if err != nil {
		return nil, err
	}

	return &userentity.RegisterCustomerResponse{
		UserID:      customer.ID,
		PhoneNumber: customer.PhoneNumber,
		Name:        customer.Name,
	}, nil
}
