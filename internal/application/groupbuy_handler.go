package application

import (
	"group-buy-market-go/internal/domain/activity/service"
	"group-buy-market-go/internal/infrastructure/dao"
)

// GroupBuyHandler handles group buy related HTTP requests
type GroupBuyHandler struct {
	activityRepo  dao.GroupBuyActivityDAO
	marketService *service.IIndexGroupBuyMarketService
}

// NewGroupBuyHandler creates a new group buy handler
func NewGroupBuyHandler(
	activityRepo dao.GroupBuyActivityDAO,
	marketService *service.IIndexGroupBuyMarketService,
) *GroupBuyHandler {
	return &GroupBuyHandler{
		activityRepo:  activityRepo,
		marketService: marketService,
	}
}
