package repository

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/dao"
	"group-buy-market-go/internal/infrastructure/po"
)

type TradeRepository struct {
	groupBuyOrderDAO     dao.GroupBuyOrderDAO
	groupBuyOrderListDAO dao.GroupBuyOrderListDAO
	groupBuyActivityDAO  dao.GroupBuyActivityDAO // 添加活动DAO
}

// NewTradeRepository creates a new trade repository
func NewTradeRepository(
	groupBuyOrderDAO dao.GroupBuyOrderDAO,
	groupBuyOrderListDAO dao.GroupBuyOrderListDAO,
	groupBuyActivityDAO dao.GroupBuyActivityDAO,
) *TradeRepository {
	return &TradeRepository{
		groupBuyOrderDAO:     groupBuyOrderDAO,
		groupBuyOrderListDAO: groupBuyOrderListDAO,
		groupBuyActivityDAO:  groupBuyActivityDAO,
	}
}

// QueryMarketPayOrderEntityByOutTradeNo queries market pay order entity by out trade number
func (r *TradeRepository) QueryMarketPayOrderEntityByOutTradeNo(ctx context.Context, userId string, outTradeNo string) (*model.MarketPayOrderEntity, error) {
	groupBuyOrderListReq := &po.GroupBuyOrderList{
		UserId:     userId,
		OutTradeNo: outTradeNo,
	}

	groupBuyOrderListRes, err := r.groupBuyOrderListDAO.QueryGroupBuyOrderRecordByOutTradeNo(ctx, groupBuyOrderListReq)
	if err != nil {
		return nil, err
	}

	if groupBuyOrderListRes == nil {
		return nil, nil
	}

	entity := &model.MarketPayOrderEntity{
		OrderId:                groupBuyOrderListRes.OrderId,
		DeductionPrice:         groupBuyOrderListRes.DeductionPrice,
		TradeOrderStatusEnumVO: model.TradeOrderStatusEnumVOValueOf(groupBuyOrderListRes.Status),
	}

	return entity, nil
}

// LockMarketPayOrder locks market pay order
func (r *TradeRepository) LockMarketPayOrder(ctx context.Context, groupBuyOrderAggregate *model.GroupBuyOrderAggregate) (*model.MarketPayOrderEntity, error) {
	// Aggregate object information
	userEntity := groupBuyOrderAggregate.UserEntity
	payActivityEntity := groupBuyOrderAggregate.PayActivityEntity
	payDiscountEntity := groupBuyOrderAggregate.PayDiscountEntity
	userTakeOrderCount := groupBuyOrderAggregate.UserTakeOrderCount // 获取用户参与订单次数

	// Check if there is a group - teamId is empty - new group, not empty - existing group
	teamId := payActivityEntity.TeamId
	if teamId == "" {
		// Generate random team ID
		teamId = generateRandomNumericString(8)

		// Build group buy order
		groupBuyOrder := &po.GroupBuyOrder{
			TeamId:         teamId,
			ActivityId:     payActivityEntity.ActivityId,
			Source:         payDiscountEntity.Source,
			Channel:        payDiscountEntity.Channel,
			OriginalPrice:  payDiscountEntity.OriginalPrice,
			DeductionPrice: payDiscountEntity.DeductionPrice,
			PayPrice:       payDiscountEntity.DeductionPrice,
			TargetCount:    payActivityEntity.TargetCount,
			CompleteCount:  0,
			LockCount:      1,
		}

		// Insert record
		err := r.groupBuyOrderDAO.Insert(ctx, groupBuyOrder)
		if err != nil {
			return nil, err
		}
	} else {
		// Update record - if update count is not 1, it means the group is full, throw an exception
		rowsAffected, err := r.groupBuyOrderDAO.UpdateAddLockCount(ctx, teamId)
		if err != nil {
			return nil, err
		}
		if rowsAffected != 1 {
			// 在Java版本中，如果更新记录不等于1，则表示拼团已满，抛出异常
			return nil, nil // Placeholder for error handling
		}
	}

	// Generate random order ID
	orderId := generateRandomNumericString(12)

	groupBuyOrderListReq := &po.GroupBuyOrderList{
		UserId:         userEntity.UserId,
		TeamId:         teamId,
		OrderId:        orderId,
		ActivityId:     payActivityEntity.ActivityId,
		StartTime:      payActivityEntity.StartTime,
		EndTime:        payActivityEntity.EndTime,
		GoodsId:        payDiscountEntity.GoodsId,
		Source:         payDiscountEntity.Source,
		Channel:        payDiscountEntity.Channel,
		OriginalPrice:  payDiscountEntity.OriginalPrice,
		DeductionPrice: payDiscountEntity.DeductionPrice,
		Status:         model.CREATE.Code(),
		OutTradeNo:     payDiscountEntity.OutTradeNo,
		// 构建 bizId 唯一值；活动id_用户id_参与次数累加
		BizId: fmt.Sprintf("%d_%s_%d", payActivityEntity.ActivityId, userEntity.UserId, userTakeOrderCount+1),
	}

	// Insert group buy order list record
	err := r.groupBuyOrderListDAO.Insert(ctx, groupBuyOrderListReq)
	if err != nil {
		// 在Java版本中处理了DuplicateKeyException
		return nil, err
	}

	entity := &model.MarketPayOrderEntity{
		OrderId:                orderId,
		DeductionPrice:         payDiscountEntity.DeductionPrice,
		TradeOrderStatusEnumVO: model.CREATE,
	}

	return entity, nil
}

// QueryGroupBuyProgress queries group buy progress
func (r *TradeRepository) QueryGroupBuyProgress(ctx context.Context, teamId string) (*model.GroupBuyProgressVO, error) {
	groupBuyOrder, err := r.groupBuyOrderDAO.QueryGroupBuyProgress(ctx, teamId)
	if err != nil {
		return nil, err
	}

	if groupBuyOrder == nil {
		return nil, nil
	}

	vo := &model.GroupBuyProgressVO{
		CompleteCount: groupBuyOrder.CompleteCount,
		TargetCount:   groupBuyOrder.TargetCount,
		LockCount:     groupBuyOrder.LockCount,
	}

	return vo, nil
}

// QueryGroupBuyActivityEntityByActivityId queries group buy activity entity by activity ID
func (r *TradeRepository) QueryGroupBuyActivityEntityByActivityId(ctx context.Context, activityId int64) (*model.GroupBuyActivityEntity, error) {
	groupBuyActivity, err := r.groupBuyActivityDAO.QueryGroupBuyActivityByActivityId(ctx, activityId)
	if err != nil {
		return nil, err
	}

	if groupBuyActivity == nil {
		return nil, nil
	}

	entity := &model.GroupBuyActivityEntity{
		ActivityId:     groupBuyActivity.ActivityId,
		ActivityName:   groupBuyActivity.ActivityName,
		DiscountId:     groupBuyActivity.DiscountId,
		GroupType:      int(groupBuyActivity.GroupType),
		TakeLimitCount: int(groupBuyActivity.TakeLimitCount),
		Target:         int(groupBuyActivity.Target),
		ValidTime:      int(groupBuyActivity.ValidTime),
		Status:         model.ActivityStatusEnumVOValueOf(int(groupBuyActivity.Status)),
		StartTime:      groupBuyActivity.StartTime,
		EndTime:        groupBuyActivity.EndTime,
		TagId:          groupBuyActivity.TagId,
		TagScope:       groupBuyActivity.TagScope,
	}

	return entity, nil
}

// QueryOrderCountByActivityId queries order count by activity ID and user ID
func (r *TradeRepository) QueryOrderCountByActivityId(ctx context.Context, activityId int64, userId string) (int, error) {
	groupBuyOrderListReq := &po.GroupBuyOrderList{
		ActivityId: activityId,
		UserId:     userId,
	}

	count, err := r.groupBuyOrderListDAO.QueryOrderCountByActivityId(ctx, groupBuyOrderListReq)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// generateRandomNumericString generates a random numeric string of specified length
func generateRandomNumericString(length int) string {
	const charset = "0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
