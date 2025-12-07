package application

import (
	"encoding/json"
	"group-buy-market-go/internal/domain"
	"group-buy-market-go/internal/infrastructure/dao"
	"net/http"
	"strconv"
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
	activities, err := h.activityRepo.QueryGroupBuyActivityList()
	if err != nil {
		http.Error(w, "Failed to retrieve activities: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}

// GetAllDiscounts retrieves all group buy discounts
func (h *GroupBuyHandler) GetAllDiscounts(w http.ResponseWriter, r *http.Request) {
	discounts, err := h.groupBuyService.GetAllDiscounts()
	if err != nil {
		http.Error(w, "Failed to retrieve discounts: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(discounts)
}

// GetDiscountByID retrieves a group buy discount by ID
func (h *GroupBuyHandler) GetDiscountByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from query parameters
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing discount ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid discount ID", http.StatusBadRequest)
		return
	}

	discount, err := h.groupBuyService.GetDiscountByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve discount: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if discount == nil {
		http.Error(w, "Discount not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(discount)
}
