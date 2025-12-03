package application

import (
	"encoding/json"
	"group-buy-market-go/internal/domain"
	"group-buy-market-go/internal/infrastructure/dao"
	"net/http"
)

// GroupBuyHandler handles group buy related HTTP requests
type GroupBuyHandler struct {
	groupBuyService *domain.GroupBuyService
	activityRepo    dao.GroupBuyActivityDAO
}

// NewGroupBuyHandler creates a new group buy handler
func NewGroupBuyHandler(groupBuyService *domain.GroupBuyService, activityRepo dao.GroupBuyActivityDAO) *GroupBuyHandler {
	return &GroupBuyHandler{
		groupBuyService: groupBuyService,
		activityRepo:    activityRepo,
	}
}

// GetAllActivities retrieves all group buy activities
func (h *GroupBuyHandler) GetAllActivities(w http.ResponseWriter, r *http.Request) {
	activities, err := h.activityRepo.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieve activities: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}
