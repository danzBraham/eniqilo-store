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
		Id:          ulid.Make().String(),
		PhoneNumber: payload.PhoneNumber,
		Name:        payload.Name,
		Password:    hashedPassword,
		Role:        userentity.Staff,
	}

	err = s.UserRepository.CreateUser(ctx, staff)
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(2*time.Hour, staff.Id)
	if err != nil {
		return nil, err
	}

	return &userentity.RegisterStaffResponse{
		UserId:      staff.Id,
		PhoneNumber: staff.PhoneNumber,
		Name:        staff.Name,
		AccessToken: token,
	}, nil
}
