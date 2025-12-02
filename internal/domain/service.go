package domain

// ProductService provides product domain services
type ProductService struct {
	repo ProductRepository
}

// NewProductService creates a new product service
func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// GetUserDiscount calculates discount for a user
func (s *ProductService) GetUserDiscount(user *User) float64 {
	// Simple discount logic - could be more complex in reality
	if user.Username == "vip" {
		return 0.1 // 10% discount for VIP users
	}
	return 0.0
}
