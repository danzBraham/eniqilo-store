package checkouterror

import "errors"

var (
	ErrCustomerIDNotFound         = errors.New("customer ID not found")
	ErrProductIDNotFound          = errors.New("product ID not found")
	ErrOneOfProductNotAvailable   = errors.New("one of product not available")
	ErrOneOfProductStockNotEnough = errors.New("one of product stock not enough")
	ErrPaidNotEnough              = errors.New("paid not enough")
	ErrChangeNotRight             = errors.New("change not right")
)
