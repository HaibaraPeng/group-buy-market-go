package repository

import (
	"context"
	"strconv"

	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/infrastructure/dao"
	"group-buy-market-go/internal/infrastructure/po"
)

type ActivityRepository struct {
	groupBuyActivityDAO dao.GroupBuyActivityDAO
	groupBuyDiscountDAO dao.GroupBuyDiscountDAO
	skuDAO              dao.SkuDAO
	scSkuActivityDAO    dao.SCSkuActivityDAO
}

// NewActivityRepository creates a new activity repository
func NewActivityRepository(groupBuyActivityDAO dao.GroupBuyActivityDAO, groupBuyDiscountDAO dao.GroupBuyDiscountDAO,
	skuDAO dao.SkuDAO, scSkuActivityDAO dao.SCSkuActivityDAO) *ActivityRepository {
	return &ActivityRepository{
		groupBuyActivityDAO: groupBuyActivityDAO,
		groupBuyDiscountDAO: groupBuyDiscountDAO,
		skuDAO:              skuDAO,
		scSkuActivityDAO:    scSkuActivityDAO,
	}
}

// QueryGroupBuyActivityDiscountVO queries group buy activity and its associated discount by source and channel
func (r *ActivityRepository) QueryGroupBuyActivityDiscountVO(ctx context.Context, activityId int64) (*model.GroupBuyActivityDiscountVO, error) {
	// Query the latest valid activity by source and channel
	groupBuyActivityRes, err := r.groupBuyActivityDAO.FindValidByActivityID(ctx, activityId)
	if err != nil {
		return nil, err
	}

	// If no activity found, return nil
	if groupBuyActivityRes == nil {
		return nil, nil
	}

	// Convert discount ID from string to int
	discountID, err := strconv.Atoi(groupBuyActivityRes.DiscountId)
	if err != nil {
		return nil, err
	}

	// Query discount by discount id
	groupBuyDiscountRes, err := r.groupBuyDiscountDAO.FindByDiscountID(ctx, discountID)
	if err != nil {
		return nil, err
	}

	// If no discount found, return nil
	if groupBuyDiscountRes == nil {
		return nil, nil
	}

	// Build the GroupBuyDiscount VO
	groupBuyDiscount := &model.GroupBuyDiscountVO{
		DiscountName: groupBuyDiscountRes.DiscountName,
		DiscountDesc: groupBuyDiscountRes.DiscountDesc,
		DiscountType: model.DiscountTypeEnum(groupBuyDiscountRes.DiscountType), // 类型转换
		MarketPlan:   model.MarketPlanEnum(groupBuyDiscountRes.MarketPlan),     // 类型转换
		MarketExpr:   groupBuyDiscountRes.MarketExpr,
		TagId:        groupBuyDiscountRes.TagId,
	}

	// Build and return the final VO
	vo := &model.GroupBuyActivityDiscountVO{
		ActivityId:       groupBuyActivityRes.ActivityId,
		ActivityName:     groupBuyActivityRes.ActivityName,
		GroupBuyDiscount: groupBuyDiscount,
		GroupType:        groupBuyActivityRes.GroupType,
		TakeLimitCount:   groupBuyActivityRes.TakeLimitCount,
		Target:           groupBuyActivityRes.Target,
		ValidTime:        groupBuyActivityRes.ValidTime,
		Status:           groupBuyActivityRes.Status,
		StartTime:        groupBuyActivityRes.StartTime,
		EndTime:          groupBuyActivityRes.EndTime,
		TagId:            groupBuyActivityRes.TagId,
		TagScope:         groupBuyActivityRes.TagScope,
	}

	return vo, nil
}

// QuerySkuByGoodsId queries sku by goods id
func (r *ActivityRepository) QuerySkuByGoodsId(goodsId string) (*model.SkuVO, error) {
	// Query sku by goods id
	sku, err := r.skuDAO.FindByGoodsId(goodsId)
	if err != nil {
		return nil, err
	}

	// If no sku found, return nil
	if sku == nil {
		return nil, nil
	}

	// Build and return the SkuVO
	skuVO := &model.SkuVO{
		GoodsId:       sku.GoodsId,
		GoodsName:     sku.GoodsName,
		OriginalPrice: sku.OriginalPrice,
	}

	return skuVO, nil
}

// QuerySCSkuActivityBySCGoodsId queries sc sku activity by source, channel and goods id
func (r *ActivityRepository) QuerySCSkuActivityBySCGoodsId(ctx context.Context, source string, channel string, goodsId string) (*model.SCSkuActivityVO, error) {
	// Create SCSkuActivity PO with given parameters
	scSkuActivityReq := &po.SCSkuActivity{
		Source:  source,
		Channel: channel,
		GoodsId: goodsId,
	}

	// Query sc sku activity by source, channel and goods id
	scSkuActivity, err := r.scSkuActivityDAO.QuerySCSkuActivityBySCGoodsId(ctx, scSkuActivityReq)
	if err != nil {
		return nil, err
	}

	// If no sc sku activity found, return nil
	if scSkuActivity == nil {
		return nil, nil
	}

	// Build and return the SCSkuActivityVO
	scSkuActivityVO := &model.SCSkuActivityVO{
		Source:     scSkuActivity.Source,
		Channel:    scSkuActivity.Channel,
		ActivityId: scSkuActivity.ActivityId,
		GoodsId:    scSkuActivity.GoodsId,
	}

	return scSkuActivityVO, nil
}
