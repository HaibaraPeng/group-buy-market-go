package repository

import (
	"context"
	"group-buy-market-go/internal/common/utils"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/infrastructure/dao"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/dcc" // 添加dcc包
	"group-buy-market-go/internal/infrastructure/po"
	"math/rand"
	"time"
)

type ActivityRepository struct {
	groupBuyActivityDAO  dao.GroupBuyActivityDAO
	groupBuyDiscountDAO  dao.GroupBuyDiscountDAO
	skuDAO               dao.SkuDAO
	scSkuActivityDAO     dao.SCSkuActivityDAO
	groupBuyOrderDAO     dao.GroupBuyOrderDAO
	groupBuyOrderListDAO dao.GroupBuyOrderListDAO
	data                 *data.Data
	dcc                  *dcc.DCC // 添加DCC
}

// NewActivityRepository creates a new activity repository
func NewActivityRepository(groupBuyActivityDAO dao.GroupBuyActivityDAO, groupBuyDiscountDAO dao.GroupBuyDiscountDAO,
	skuDAO dao.SkuDAO, scSkuActivityDAO dao.SCSkuActivityDAO, groupBuyOrderDAO dao.GroupBuyOrderDAO,
	groupBuyOrderListDAO dao.GroupBuyOrderListDAO, data *data.Data, dcc *dcc.DCC) *ActivityRepository { // 添加新参数
	return &ActivityRepository{
		groupBuyActivityDAO:  groupBuyActivityDAO,
		groupBuyDiscountDAO:  groupBuyDiscountDAO,
		skuDAO:               skuDAO,
		scSkuActivityDAO:     scSkuActivityDAO,
		groupBuyOrderDAO:     groupBuyOrderDAO,
		groupBuyOrderListDAO: groupBuyOrderListDAO,
		data:                 data,
		dcc:                  dcc, // 初始化DCC服务
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
func (r *ActivityRepository) QuerySkuByGoodsId(ctx context.Context, goodsId string) (*model.SkuVO, error) {
	// Query sku by goods id
	sku, err := r.skuDAO.FindByGoodsId(ctx, goodsId)
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
	exists, err := r.data.Rdb(ctx).Exists(ctx, tagId).Result()
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
	bitValue, err := r.data.Rdb(ctx).GetBit(ctx, tagId, userIndex).Result()
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

// QueryInProgressUserGroupBuyOrderDetailListByOwner 查询拥有者参与的拼团详情列表
func (r *ActivityRepository) QueryInProgressUserGroupBuyOrderDetailListByOwner(ctx context.Context, activityId int64, userId string, ownerCount int) ([]*model.UserGroupBuyOrderDetailEntity, error) {
	// 1. 根据用户ID、活动ID，查询用户参与的拼团队伍
	req := &po.GroupBuyOrderList{
		ActivityId: activityId,
		UserId:     userId,
		Count:      int64(ownerCount),
	}
	groupBuyOrderLists, err := r.groupBuyOrderListDAO.QueryInProgressUserGroupBuyOrderDetailListByUserId(ctx, req)
	if err != nil {
		return nil, err
	}
	if groupBuyOrderLists == nil || len(groupBuyOrderLists) == 0 {
		return nil, nil
	}

	// 2. 过滤队伍获取 TeamId
	teamIds := make([]string, 0)
	for _, list := range groupBuyOrderLists {
		if list.TeamId != "" {
			teamIds = append(teamIds, list.TeamId)
		}
	}

	// 3. 查询队伍明细，组装Map结构
	groupBuyOrders, err := r.groupBuyOrderDAO.QueryGroupBuyProgressByTeamIds(ctx, teamIds)
	if err != nil {
		return nil, err
	}
	if groupBuyOrders == nil || len(groupBuyOrders) == 0 {
		return nil, nil
	}

	groupBuyOrderMap := make(map[string]*po.GroupBuyOrder)
	for _, order := range groupBuyOrders {
		groupBuyOrderMap[order.TeamId] = order
	}

	// 4. 转换数据
	userGroupBuyOrderDetailEntities := make([]*model.UserGroupBuyOrderDetailEntity, 0)
	for _, list := range groupBuyOrderLists {
		teamId := list.TeamId
		groupBuyOrder, exists := groupBuyOrderMap[teamId]
		if !exists {
			continue
		}

		userGroupBuyOrderDetailEntity := &model.UserGroupBuyOrderDetailEntity{
			UserId:         list.UserId,
			TeamId:         groupBuyOrder.TeamId,
			ActivityId:     groupBuyOrder.ActivityId,
			TargetCount:    groupBuyOrder.TargetCount,
			CompleteCount:  groupBuyOrder.CompleteCount,
			LockCount:      groupBuyOrder.LockCount,
			ValidStartTime: groupBuyOrder.ValidStartTime,
			ValidEndTime:   groupBuyOrder.ValidEndTime,
			OutTradeNo:     list.OutTradeNo,
		}

		userGroupBuyOrderDetailEntities = append(userGroupBuyOrderDetailEntities, userGroupBuyOrderDetailEntity)
	}

	return userGroupBuyOrderDetailEntities, nil
}

// QueryInProgressUserGroupBuyOrderDetailListByRandom 随机查询拼团详情列表
func (r *ActivityRepository) QueryInProgressUserGroupBuyOrderDetailListByRandom(ctx context.Context, activityId int64, userId string, randomCount int) ([]*model.UserGroupBuyOrderDetailEntity, error) {
	// 1. 根据用户ID、活动ID，查询用户参与的拼团队伍
	req := &po.GroupBuyOrderList{
		ActivityId: activityId,
		UserId:     userId,
		Count:      int64(randomCount * 2), // 查询2倍的量，之后其中 randomCount 数量
	}
	groupBuyOrderLists, err := r.groupBuyOrderListDAO.QueryInProgressUserGroupBuyOrderDetailListByRandom(ctx, req)
	if err != nil {
		return nil, err
	}
	if groupBuyOrderLists == nil || len(groupBuyOrderLists) == 0 {
		return nil, nil
	}

	// 判断总量是否大于 randomCount
	if len(groupBuyOrderLists) > randomCount {
		// 随机打乱列表
		rand.Seed(time.Now().UnixNano())
		for i := len(groupBuyOrderLists) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			groupBuyOrderLists[i], groupBuyOrderLists[j] = groupBuyOrderLists[j], groupBuyOrderLists[i]
		}
		// 获取前 randomCount 个元素
		groupBuyOrderLists = groupBuyOrderLists[:randomCount]
	}

	// 2. 过滤队伍获取 TeamId
	teamIds := make([]string, 0)
	for _, list := range groupBuyOrderLists {
		if list.TeamId != "" {
			teamIds = append(teamIds, list.TeamId)
		}
	}

	// 3. 查询队伍明细，组装Map结构
	groupBuyOrders, err := r.groupBuyOrderDAO.QueryGroupBuyProgressByTeamIds(ctx, teamIds)
	if err != nil {
		return nil, err
	}
	if groupBuyOrders == nil || len(groupBuyOrders) == 0 {
		return nil, nil
	}

	groupBuyOrderMap := make(map[string]*po.GroupBuyOrder)
	for _, order := range groupBuyOrders {
		groupBuyOrderMap[order.TeamId] = order
	}

	// 4. 转换数据
	userGroupBuyOrderDetailEntities := make([]*model.UserGroupBuyOrderDetailEntity, 0)
	for _, list := range groupBuyOrderLists {
		teamId := list.TeamId
		groupBuyOrder, exists := groupBuyOrderMap[teamId]
		if !exists {
			continue
		}

		userGroupBuyOrderDetailEntity := &model.UserGroupBuyOrderDetailEntity{
			UserId:         list.UserId,
			TeamId:         groupBuyOrder.TeamId,
			ActivityId:     groupBuyOrder.ActivityId,
			TargetCount:    groupBuyOrder.TargetCount,
			CompleteCount:  groupBuyOrder.CompleteCount,
			LockCount:      groupBuyOrder.LockCount,
			ValidStartTime: groupBuyOrder.ValidStartTime,
			ValidEndTime:   groupBuyOrder.ValidEndTime,
		}

		userGroupBuyOrderDetailEntities = append(userGroupBuyOrderDetailEntities, userGroupBuyOrderDetailEntity)
	}

	return userGroupBuyOrderDetailEntities, nil
}

// QueryTeamStatisticByActivityId 根据活动ID查询团队统计信息
func (r *ActivityRepository) QueryTeamStatisticByActivityId(ctx context.Context, activityId int64) (*model.TeamStatisticVO, error) {
	// 1. 根据活动ID查询拼团队伍
	groupBuyOrderLists, err := r.groupBuyOrderListDAO.QueryInProgressUserGroupBuyOrderDetailListByActivityId(ctx, activityId)
	if err != nil {
		return nil, err
	}

	if groupBuyOrderLists == nil || len(groupBuyOrderLists) == 0 {
		return &model.TeamStatisticVO{
			AllTeamCount:         0,
			AllTeamCompleteCount: 0,
			AllTeamUserCount:     0,
		}, nil
	}

	// 2. 过滤队伍获取 TeamId
	teamIds := make([]string, 0)
	for _, list := range groupBuyOrderLists {
		if list.TeamId != "" {
			teamIds = append(teamIds, list.TeamId)
		}
	}

	// 3. 统计数据
	allTeamCount, err := r.groupBuyOrderDAO.QueryAllTeamCount(ctx, teamIds)
	if err != nil {
		return nil, err
	}

	allTeamCompleteCount, err := r.groupBuyOrderDAO.QueryAllTeamCompleteCount(ctx, teamIds)
	if err != nil {
		return nil, err
	}

	allTeamUserCount, err := r.groupBuyOrderDAO.QueryAllUserCount(ctx, teamIds)
	if err != nil {
		return nil, err
	}

	// 4. 构建对象
	return &model.TeamStatisticVO{
		AllTeamCount:         allTeamCount,
		AllTeamCompleteCount: allTeamCompleteCount,
		AllTeamUserCount:     allTeamUserCount,
	}, nil
}
