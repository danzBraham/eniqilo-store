package productentity

type Category string

var (
	Clothing    Category = "Clothing"
	Accessories Category = "Accessories"
	Footwear    Category = "Footwear"
	Beverages   Category = "Beverages"
)

type Product struct {
	ID          string
	Name        string
	SKU         string
	Category    Category
	ImageURL    string
	Notes       string
	Price       int
	Stock       int
	Location    string
	IsAvailable bool
	CreatedAt   string
	UpdatedAt   string
}

type CreateProductRequest struct {
	Name        string   `json:"name" validate:"required,min=1,max=30"`
	SKU         string   `json:"sku" validate:"required,min=1,max=30"`
	Category    Category `json:"category" validate:"required,oneof='Clothing' 'Accessories' 'Footwear' 'Beverages'"`
	ImageURL    string   `json:"imageUrl" validate:"required,http_url"`
	Notes       string   `json:"notes" validate:"required,min=1,max=200"`
	Price       int      `json:"price" validate:"required,min=1"`
	Stock       int      `json:"stock" validate:"required,min=0,max=100000"`
	Location    string   `json:"location" validate:"required,min=1,max=200"`
	IsAvailable bool     `json:"isAvailable" validate:"required"`
}

type CreateProductResponse struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}
