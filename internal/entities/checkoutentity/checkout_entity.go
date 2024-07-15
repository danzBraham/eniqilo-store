package checkoutentity

type Transaction struct {
	ID         string
	CustomerID string
	TotalPrice int
	Paid       int
	Change     int
	CreatedAt  string
	UpdatedAt  string
}

type Checkout struct {
	ID            string
	TransactionID string
	ProductID     string
	Quantity      int
	Price         int
}

type ProductDetails struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

type CheckoutProductRequest struct {
	CustomerID     string           `json:"customerId"`
	ProductDetails []ProductDetails `json:"productDetails" validate:"required,min=1,dive"`
	Paid           int              `json:"paid" validate:"required,min=1"`
	Change         int              `json:"change" validate:"number,min=0"`
}

type CheckoutHistoryQueryParams struct {
	CustomerID string
	Limit      int
	Offset     int
	CreatedAt  string
}

type GetCheckoutHistoryResponse struct {
	TransactionID  string            `json:"transactionId"`
	CustomerID     string            `json:"customerId"`
	ProductDetails []*ProductDetails `json:"productDetails"`
	Paid           int               `json:"paid"`
	Change         int               `json:"change"`
	CreatedAt      string            `json:"createdAt"`
}
