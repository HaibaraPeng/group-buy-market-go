package infrastructure

import (
	"group-buy-market-go/internal/domain"
	"gorm.io/gorm"
)

// MySQLProductRepository is a GORM implementation of ProductRepository
type MySQLProductRepository struct {
	db *gorm.DB
}

// NewMySQLProductRepository creates a new MySQL product repository
func NewMySQLProductRepository(db *gorm.DB) *MySQLProductRepository {
	return &MySQLProductRepository{
		db: db,
	}
}

// Save saves a product
func (r *MySQLProductRepository) Save(product *domain.Product) error {
	return r.db.Save(product).Error
}

// FindByID finds a product by ID
func (r *MySQLProductRepository) FindByID(id int64) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

// FindAll returns all products
func (r *MySQLProductRepository) FindAll() ([]*domain.Product, error) {
	var products []*domain.Product
	err := r.db.Find(&products).Error
	return products, err
}