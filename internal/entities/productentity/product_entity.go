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
	ImageURL    string   `json:"imageUrl" validate:"required,image_url"`
	Notes       string   `json:"notes" validate:"required,min=1,max=200"`
	Price       int      `json:"price" validate:"required,min=1"`
	Stock       int      `json:"stock" validate:"required,min=0,max=100000"`
	Location    string   `json:"location" validate:"required,min=1,max=200"`
	IsAvailable *bool    `json:"isAvailable" validate:"boolean"`
}

type CreateProductResponse struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type ProductQueryParams struct {
	ID          string
	Limit       int
	Offset      int
	Name        string
	SKU         string
	Category    Category
	Price       string
	InStock     string
	IsAvailable string
	CreatedAt   string
}

type GetProductResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	SKU         string   `json:"sku"`
	Category    Category `json:"category"`
	ImageURL    string   `json:"imageUrl"`
	Notes       string   `json:"notes"`
	Price       int      `json:"price"`
	Stock       int      `json:"stock"`
	Location    string   `json:"location"`
	IsAvailable bool     `json:"isAvailable"`
	CreatedAt   string   `json:"createdAt"`
}

type UpdateProductRequest struct {
	Name        string   `json:"name" validate:"required,min=1,max=30"`
	SKU         string   `json:"sku" validate:"required,min=1,max=30"`
	Category    Category `json:"category" validate:"required,oneof='Clothing' 'Accessories' 'Footwear' 'Beverages'"`
	ImageURL    string   `json:"imageUrl" validate:"required,image_url"`
	Notes       string   `json:"notes" validate:"required,min=1,max=200"`
	Price       int      `json:"price" validate:"required,min=1"`
	Stock       int      `json:"stock" validate:"required,min=0,max=100000"`
	Location    string   `json:"location" validate:"required,min=1,max=200"`
	IsAvailable *bool    `json:"isAvailable" validate:"boolean"`
}
