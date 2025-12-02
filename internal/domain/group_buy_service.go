package domain

// GroupBuyService provides group buying domain services
type GroupBuyService struct {
	repo GroupBuyActivityRepository
}

// NewGroupBuyService creates a new group buy service
func NewGroupBuyService(repo GroupBuyActivityRepository) *GroupBuyService {
	return &GroupBuyService{repo: repo}
}

// IsValid checks if a group buy activity is valid
func (s *GroupBuyService) IsValid(activity *GroupBuyActivity) bool {
	// Check if activity is in active status
	if activity.Status != 1 {
		return false
	}

	// Could add more validation logic here
	return true
}

// CanJoin checks if a user can join a group buy activity
func (s *GroupBuyService) CanJoin(activity *GroupBuyActivity, userID int64) bool {
	// Check if activity is valid
	if !s.IsValid(activity) {
		return false
	}

	// Could add more business logic here
	return true
}
