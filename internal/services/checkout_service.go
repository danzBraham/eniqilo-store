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
	checkoutHistories, err := s.CheckoutRepository.GetCheckoutHistories(ctx, params)
	if err != nil {
		return nil, err
	}

	checkoutHistoryResponses := make([]*checkoutentity.GetCheckoutHistoryResponse, 0, params.Limit)
	for _, checkoutHistory := range checkoutHistories {
		checkoutProducts, err := s.CheckoutRepository.GetCheckoutProducts(ctx, checkoutHistory.ID)
		if err != nil {
			return nil, err
		}

		productDetails := make([]*checkoutentity.ProductDetails, 0, len(checkoutProducts))
		for _, checkoutProduct := range checkoutProducts {
			productDetails = append(productDetails, &checkoutentity.ProductDetails{
				ProductID: checkoutProduct.ProductID,
				Quantity:  checkoutProduct.Quantity,
			})
		}

		checkoutHistoryResponses = append(checkoutHistoryResponses, &checkoutentity.GetCheckoutHistoryResponse{
			TransactionID:  checkoutHistory.ID,
			CustomerID:     checkoutHistory.CustomerID,
			ProductDetails: productDetails,
			Paid:           checkoutHistory.Paid,
			Change:         checkoutHistory.Change,
			CreatedAt:      checkoutHistory.CreatedAt,
		})
	}

	return checkoutHistoryResponses, nil
}
