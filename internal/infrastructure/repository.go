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

// InMemoryGroupBuyActivityRepository is an in-memory implementation of GroupBuyActivityRepository
type InMemoryGroupBuyActivityRepository struct {
	activities map[int64]*domain.GroupBuyActivity
}

// NewInMemoryGroupBuyActivityRepository creates a new in-memory group buy activity repository
func NewInMemoryGroupBuyActivityRepository() *InMemoryGroupBuyActivityRepository {
	return &InMemoryGroupBuyActivityRepository{
		activities: make(map[int64]*domain.GroupBuyActivity),
	}
}

// Save saves a group buy activity
func (r *InMemoryGroupBuyActivityRepository) Save(activity *domain.GroupBuyActivity) error {
	r.activities[activity.ID] = activity
	return nil
}

// FindByID finds a group buy activity by ID
func (r *InMemoryGroupBuyActivityRepository) FindByID(id int64) (*domain.GroupBuyActivity, error) {
	activity, exists := r.activities[id]
	if !exists {
		return nil, nil
	}
	return activity, nil
}

// FindByActivityID finds a group buy activity by activity ID
func (r *InMemoryGroupBuyActivityRepository) FindByActivityID(activityID int64) (*domain.GroupBuyActivity, error) {
	for _, activity := range r.activities {
		if activity.ActivityId == activityID {
			return activity, nil
		}
	}
	return nil, nil
}

// FindAll returns all group buy activities
func (r *InMemoryGroupBuyActivityRepository) FindAll() ([]*domain.GroupBuyActivity, error) {
	var activities []*domain.GroupBuyActivity
	for _, activity := range r.activities {
		activities = append(activities, activity)
	}
	return activities, nil
}

// UpdateStatus updates the status of a group buy activity
func (r *InMemoryGroupBuyActivityRepository) UpdateStatus(id int64, status int) error {
	activity, exists := r.activities[id]
	if !exists {
		return nil
	}
	activity.Status = status
	return nil
}
