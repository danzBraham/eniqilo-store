package services

import (
	"context"

	"github.com/danzBraham/eniqilo-store/internal/entities/checkoutentity"
	"github.com/danzBraham/eniqilo-store/internal/entities/userentity"
	"github.com/danzBraham/eniqilo-store/internal/errors/checkouterror"
	"github.com/danzBraham/eniqilo-store/internal/repositories"
	"github.com/oklog/ulid/v2"
)

type CheckoutService interface {
	CheckoutProduct(ctx context.Context, payload *checkoutentity.CheckoutProductRequest) error
	GetCheckoutHistories(ctx context.Context, params *checkoutentity.CheckoutHistoryQueryParams) ([]*checkoutentity.GetCheckoutHistoryResponse, error)
}

type CheckoutServiceImpl struct {
	CheckoutRepository repositories.CheckoutRepository
	UserRepository     repositories.UserRepository
}

func NewCheckoutService(
	checkoutRepository repositories.CheckoutRepository,
	userRepository repositories.UserRepository,
) CheckoutService {
	return &CheckoutServiceImpl{
		CheckoutRepository: checkoutRepository,
		UserRepository:     userRepository,
	}
}

func (s *CheckoutServiceImpl) CheckoutProduct(ctx context.Context, payload *checkoutentity.CheckoutProductRequest) error {
	isCustomerIDExists, err := s.UserRepository.IsUserIDByItsRoleExists(ctx, payload.CustomerID, userentity.Customer)
	if err != nil {
		return err
	}
	if !isCustomerIDExists {
		return checkouterror.ErrCustomerIDNotFound
	}

	transaction := &checkoutentity.Transaction{
		ID:         ulid.Make().String(),
		CustomerID: payload.CustomerID,
	}

	err = s.CheckoutRepository.CreateCheckoutProduct(ctx, transaction, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *CheckoutServiceImpl) GetCheckoutHistories(ctx context.Context, params *checkoutentity.CheckoutHistoryQueryParams) ([]*checkoutentity.GetCheckoutHistoryResponse, error) {
	return s.CheckoutRepository.GetCheckoutHistories(ctx, params)
}
