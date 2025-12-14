package application

import (
	"encoding/json"
	"group-buy-market-go/internal/domain"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service"
	"group-buy-market-go/internal/infrastructure/dao"
	"net/http"
)

// GroupBuyHandler handles group buy related HTTP requests
type GroupBuyHandler struct {
	groupBuyService *domain.GroupBuyService
	activityRepo    dao.GroupBuyActivityDAO
	marketService   *service.IIndexGroupBuyMarketService
}

// NewGroupBuyHandler creates a new group buy handler
func NewGroupBuyHandler(
	groupBuyService *domain.GroupBuyService,
	activityRepo dao.GroupBuyActivityDAO,
	marketService *service.IIndexGroupBuyMarketService,
) *GroupBuyHandler {
	return &GroupBuyHandler{
		groupBuyService: groupBuyService,
		activityRepo:    activityRepo,
		marketService:   marketService,
	}
}

// MarketTrial handles market trial requests
func (h *GroupBuyHandler) MarketTrial(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	userId := r.URL.Query().Get("userId")
	goodsId := r.URL.Query().Get("goodsId")
	source := r.URL.Query().Get("source")
	channel := r.URL.Query().Get("channel")

	// Validate required parameters
	if userId == "" || goodsId == "" || source == "" || channel == "" {
		http.Error(w, "Missing required parameters: userId, goodsId, source, channel", http.StatusBadRequest)
		return
	}

	// Create market product entity
	marketProduct := &model.MarketProductEntity{
		UserId:  userId,
		GoodsId: goodsId,
		Source:  source,
		Channel: channel,
	}

	// Call the market service
	trialResult, err := h.marketService.IndexMarketTrial(marketProduct)
	if err != nil {
		http.Error(w, "Failed to process market trial: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trialResult)
}
