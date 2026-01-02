package repository

import (
	"context"
	"group-buy-market-go/internal/common/utils"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/infrastructure"
	"group-buy-market-go/internal/infrastructure/dao"
	"group-buy-market-go/internal/infrastructure/dcc" // 添加dcc包
	"group-buy-market-go/internal/infrastructure/po"
)

type ActivityRepository struct {
	groupBuyActivityDAO dao.GroupBuyActivityDAO
	groupBuyDiscountDAO dao.GroupBuyDiscountDAO
	skuDAO              dao.SkuDAO
	scSkuActivityDAO    dao.SCSkuActivityDAO
	data                infrastructure.Data
	dcc                 *dcc.DCC // 添加DCC
}

// NewActivityRepository creates a new activity repository
func NewActivityRepository(groupBuyActivityDAO dao.GroupBuyActivityDAO, groupBuyDiscountDAO dao.GroupBuyDiscountDAO,
	skuDAO dao.SkuDAO, scSkuActivityDAO dao.SCSkuActivityDAO, data infrastructure.Data, dcc *dcc.DCC) *ActivityRepository { // 添加dcc参数
	return &ActivityRepository{
		groupBuyActivityDAO: groupBuyActivityDAO,
		groupBuyDiscountDAO: groupBuyDiscountDAO,
		skuDAO:              skuDAO,
		scSkuActivityDAO:    scSkuActivityDAO,
		data:                data,
		dcc:                 dcc, // 初始化DCC服务
	}
}

// QueryGroupBuyActivityDiscountVO queries group buy activity and its associated discount by activity id
func (r *ActivityRepository) QueryGroupBuyActivityDiscountVO(ctx context.Context, activityId int64) (*model.GroupBuyActivityDiscountVO, error) {
	// Query the activity by activity id
	groupBuyActivityRes, err := r.groupBuyActivityDAO.FindValidByActivityID(ctx, activityId)
	if err != nil {
		return nil, err
	}

	// If no activity found, return nil
	if groupBuyActivityRes == nil {
		return nil, nil
	}

	// Get discount ID from activity
	discountId := groupBuyActivityRes.DiscountId

	// Query discount by discount id using the method that matches Java implementation
	groupBuyDiscountRes, err := r.groupBuyDiscountDAO.QueryGroupBuyActivityDiscountByDiscountId(ctx, discountId)
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

// IsTagCrowdRange checks if user is in the tag crowd range
func (r *ActivityRepository) IsTagCrowdRange(ctx context.Context, tagId string, userId string) (bool, error) {
	// Check if bitset exists for tagId
	exists, err := r.data.Rdb.Exists(ctx, tagId).Result()
	if err != nil {
		return false, err
	}

	// If bitset doesn't exist, return true
	if exists == 0 {
		return true, nil
	}

	// Calculate user index (simplified - in real implementation this would be more complex)
	// This is a placeholder implementation - you'll need to adjust based on your actual userId to index mapping
	userIndex := utils.GetIndexFromUserId(userId)

	// Get bit value at user index
	bitValue, err := r.data.Rdb.GetBit(ctx, tagId, userIndex).Result()
	if err != nil {
		return false, err
	}

	// Return true if bit is set (user is in the crowd)
	return bitValue == 1, nil
}

// DowngradeSwitch 判断是否开启降级开关
func (r *ActivityRepository) DowngradeSwitch() bool {
	return r.dcc.IsDowngradeSwitch()
}

// CutRange 判断用户是否在切量范围内
func (r *ActivityRepository) CutRange(userId string) (bool, error) {
	return r.dcc.IsCutRange(userId)
}
