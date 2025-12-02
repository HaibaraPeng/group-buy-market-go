package infrastructure

import (
	"group-buy-market-go/internal/domain"
)

// InMemoryProductRepository is an in-memory implementation of ProductRepository
type InMemoryProductRepository struct {
	products map[int64]*domain.Product
}

// NewInMemoryProductRepository creates a new in-memory product repository
func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[int64]*domain.Product),
	}
}

// Save saves a product
func (r *InMemoryProductRepository) Save(product *domain.Product) error {
	r.products[product.ID] = product
	return nil
}

// FindByID finds a product by ID
func (r *InMemoryProductRepository) FindByID(id int64) (*domain.Product, error) {
	product, exists := r.products[id]
	if !exists {
		return nil, nil
	}
	return product, nil
}

// FindAll returns all products
func (r *InMemoryProductRepository) FindAll() ([]*domain.Product, error) {
	var products []*domain.Product
	for _, product := range r.products {
		products = append(products, product)
	}
	return products, nil
}
