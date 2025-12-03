package domain

import (
	"group-buy-market-go/internal/infrastructure/dao"
	"group-buy-market-go/internal/infrastructure/po"
)

// GroupBuyService provides group buying domain services
type GroupBuyService struct {
	activityRepo dao.GroupBuyActivityDAO
	discountRepo dao.GroupBuyDiscountDAO
}

// NewGroupBuyService creates a new group buy service
func NewGroupBuyService(activityRepo dao.GroupBuyActivityDAO, discountRepo dao.GroupBuyDiscountDAO) *GroupBuyService {
	return &GroupBuyService{
		activityRepo: activityRepo,
		discountRepo: discountRepo,
	}
}

// IsValid checks if a group buy activity is valid
func (s *GroupBuyService) IsValid(activity *po.GroupBuyActivity) bool {
	// Check if activity is in active status
	if activity.Status != 1 {
		return false
	}

	// Could add more validation logic here
	return true
}

// CanJoin checks if a user can join a group buy activity
func (s *GroupBuyService) CanJoin(activity *po.GroupBuyActivity, userID int64) bool {
	// Check if activity is valid
	if !s.IsValid(activity) {
		return false
	}

	// Could add more business logic here
	return true
}

// GetDiscountByID retrieves a discount by its ID
func (s *GroupBuyService) GetDiscountByID(id int64) (*po.GroupBuyDiscount, error) {
	return s.discountRepo.FindByID(id)
}

// GetAllDiscounts retrieves all discounts
func (s *GroupBuyService) GetAllDiscounts() ([]*po.GroupBuyDiscount, error) {
	return s.discountRepo.FindAll()
}
