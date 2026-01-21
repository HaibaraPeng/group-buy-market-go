package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"group-buy-market-go/internal/conf"
	"group-buy-market-go/internal/infrastructure/data"
	"math/rand"
	"time"

	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/dao"
	"group-buy-market-go/internal/infrastructure/dcc"
	"group-buy-market-go/internal/infrastructure/po"
)

type TradeRepository struct {
	data                 *data.Data
	groupBuyOrderDAO     dao.GroupBuyOrderDAO
	groupBuyOrderListDAO dao.GroupBuyOrderListDAO
	groupBuyActivityDAO  dao.GroupBuyActivityDAO // 添加活动DAO
	notifyTaskDAO        dao.NotifyTaskDAO       // 添加通知任务DAO
	dcc                  *dcc.DCC                // 添加DCC服务
	config               *conf.Data
}

// NewTradeRepository creates a new trade repository
func NewTradeRepository(
	data *data.Data,
	groupBuyOrderDAO dao.GroupBuyOrderDAO,
	groupBuyOrderListDAO dao.GroupBuyOrderListDAO,
	groupBuyActivityDAO dao.GroupBuyActivityDAO,
	notifyTaskDAO dao.NotifyTaskDAO,
	dcc *dcc.DCC, // 添加DCC服务
	config *conf.Data,
) *TradeRepository {
	return &TradeRepository{
		data:                 data,
		groupBuyOrderDAO:     groupBuyOrderDAO,
		groupBuyOrderListDAO: groupBuyOrderListDAO,
		groupBuyActivityDAO:  groupBuyActivityDAO,
		notifyTaskDAO:        notifyTaskDAO,
		dcc:                  dcc, // 初始化DCC服务
		config:               config,
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
		TeamId:                 groupBuyOrderListRes.TeamId,
		OrderId:                groupBuyOrderListRes.OrderId,
		OriginalPrice:          groupBuyOrderListRes.OriginalPrice,
		DeductionPrice:         groupBuyOrderListRes.DeductionPrice,
		PayPrice:               groupBuyOrderListRes.PayPrice,
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
	notifyConfigVO := payDiscountEntity.NotifyConfigVO
	userTakeOrderCount := groupBuyOrderAggregate.UserTakeOrderCount // 获取用户参与订单次数

	// Check if there is a group - teamId is empty - new group, not empty - existing group
	teamId := payActivityEntity.TeamId
	if teamId == "" {
		// Generate random team ID
		teamId = generateRandomNumericString(8)

		// Calculate valid start and end time
		currentTime := time.Now()
		validEndTime := currentTime.Add(time.Duration(payActivityEntity.ValidTime) * time.Minute)

		// Build group buy order
		groupBuyOrder := &po.GroupBuyOrder{
			TeamId:         teamId,
			ActivityId:     payActivityEntity.ActivityId,
			Source:         payDiscountEntity.Source,
			Channel:        payDiscountEntity.Channel,
			OriginalPrice:  payDiscountEntity.OriginalPrice,
			DeductionPrice: payDiscountEntity.DeductionPrice,
			PayPrice:       payDiscountEntity.PayPrice,
			TargetCount:    payActivityEntity.TargetCount,
			CompleteCount:  0,
			LockCount:      1,
			ValidStartTime: currentTime,
			ValidEndTime:   validEndTime,
			NotifyType:     notifyConfigVO.NotifyType.String(),
			NotifyUrl:      notifyConfigVO.NotifyUrl, // 使用notifyConfigVO中的NotifyUrl
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
			return nil, fmt.Errorf("group is full or update failed")
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
		PayPrice:       payDiscountEntity.PayPrice,
		Status:         model.CREATE.Code(),
		OutTradeNo:     payDiscountEntity.OutTradeNo,
		// 构建 bizId 唯一值；活动id_用户id_参与次数累加
		BizId: fmt.Sprintf("%d_%s_%d", payActivityEntity.ActivityId, userEntity.UserId, userTakeOrderCount+1),
	}

	// Insert group buy order list record
	err := r.groupBuyOrderListDAO.Insert(ctx, groupBuyOrderListReq)
	if err != nil {
		// 在Go版本中处理重复键异常
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
		GroupType:      groupBuyActivity.GroupType,
		TakeLimitCount: groupBuyActivity.TakeLimitCount,
		Target:         groupBuyActivity.Target,
		ValidTime:      groupBuyActivity.ValidTime,
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

// QueryGroupBuyTeamByTeamId queries group buy team by team ID
func (r *TradeRepository) QueryGroupBuyTeamByTeamId(ctx context.Context, teamId string) (*model.GroupBuyTeamEntity, error) {
	groupBuyOrder, err := r.groupBuyOrderDAO.QueryGroupBuyTeamByTeamId(ctx, teamId)
	if err != nil {
		return nil, err
	}

	if groupBuyOrder == nil {
		return nil, nil
	}

	// 处理可能为空的 NotifyType，默认为 HTTP
	notifyType := groupBuyOrder.NotifyType
	if notifyType == "" {
		notifyType = "HTTP"
	}

	// 创建 NotifyConfigVO 对象
	notifyConfigVO := &model.NotifyConfigVO{
		NotifyType: model.NotifyTypeEnumVOValueOf(notifyType),
		NotifyUrl:  groupBuyOrder.NotifyUrl,
		NotifyMQ:   r.config.Rabbitmq.Producer.TopicTeamSuccess.RoutingKey,
	}

	entity := &model.GroupBuyTeamEntity{
		TeamId:         groupBuyOrder.TeamId,
		ActivityId:     groupBuyOrder.ActivityId,
		TargetCount:    groupBuyOrder.TargetCount,
		CompleteCount:  groupBuyOrder.CompleteCount,
		LockCount:      groupBuyOrder.LockCount,
		Status:         model.GroupBuyOrderEnumVOValueOf(groupBuyOrder.Status),
		ValidStartTime: groupBuyOrder.ValidStartTime,
		ValidEndTime:   groupBuyOrder.ValidEndTime,
		NotifyConfigVO: notifyConfigVO,
	}

	return entity, nil
}

// SettlementMarketPayOrder settles market pay order
func (r *TradeRepository) SettlementMarketPayOrder(ctx context.Context, groupBuyTeamSettlementAggregate *model.GroupBuyTeamSettlementAggregate) (bool, error) {
	userEntity := groupBuyTeamSettlementAggregate.UserEntity
	groupBuyTeamEntity := groupBuyTeamSettlementAggregate.GroupBuyTeamEntity
	tradePaySuccessEntity := groupBuyTeamSettlementAggregate.TradePaySuccessEntity

	// 1. Update order list status to complete
	groupBuyOrderListReq := &po.GroupBuyOrderList{
		UserId:       userEntity.UserId,
		OutTradeNo:   tradePaySuccessEntity.OutTradeNo,
		OutTradeTime: &tradePaySuccessEntity.OutTradeTime,
	}
	err := r.data.InTx(ctx, func(ctx context.Context) error {
		rowsAffected, err := r.groupBuyOrderListDAO.UpdateOrderStatus2COMPLETE(ctx, groupBuyOrderListReq)
		if err != nil {
			return err
		}
		if rowsAffected != 1 {
			return fmt.Errorf("failed to update order status, affected rows: %d", rowsAffected)
		}

		// 2. Update complete count
		rowsAffected, err = r.groupBuyOrderDAO.UpdateAddCompleteCount(ctx, groupBuyTeamEntity.TeamId)
		if err != nil {
			return err
		}
		if rowsAffected != 1 {
			return fmt.Errorf("failed to update complete count, affected rows: %d", rowsAffected)
		}

		// 3. Update order status to complete if target is reached
		if groupBuyTeamEntity.TargetCount-groupBuyTeamEntity.CompleteCount == 1 {
			rowsAffected, err := r.groupBuyOrderDAO.UpdateOrderStatus2COMPLETE(ctx, groupBuyTeamEntity.TeamId)
			if err != nil {
				return err
			}
			if rowsAffected != 1 {
				return fmt.Errorf("failed to update order status to complete, affected rows: %d", rowsAffected)
			}

			// Query completed order transaction numbers by team ID
			outTradeNoList, err := r.groupBuyOrderListDAO.QueryGroupBuyCompleteOrderOutTradeNoListByTeamId(ctx, groupBuyTeamEntity.TeamId)
			if err != nil {
				return err
			}

			// Create notify task when group buy is completed
			notifyTask := &po.NotifyTask{
				ActivityId:    groupBuyTeamEntity.ActivityId,
				TeamId:        groupBuyTeamEntity.TeamId,
				NotifyUrl:     groupBuyTeamEntity.NotifyConfigVO.NotifyUrl,
				NotifyCount:   0,
				NotifyStatus:  0,
				ParameterJson: "",
			}

			// Prepare parameter JSON
			params := map[string]interface{}{
				"teamId":         groupBuyTeamEntity.TeamId,
				"outTradeNoList": outTradeNoList,
			}
			jsonBytes, err := json.Marshal(params)
			if err != nil {
				return err
			}
			notifyTask.ParameterJson = string(jsonBytes)

			// Insert notify task
			err = r.notifyTaskDAO.Insert(ctx, notifyTask)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	if groupBuyTeamEntity.TargetCount-groupBuyTeamEntity.CompleteCount == 1 {
		return true, nil
	}
	return false, nil
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

// IsSCBlackIntercept 判断黑名单拦截渠道，true 拦截、false 放行
func (r *TradeRepository) IsSCBlackIntercept(source, channel string) bool {
	return r.dcc.IsSCBlackIntercept(source, channel)
}

// QueryUnExecutedNotifyTaskList 查询未执行的通知任务列表
func (r *TradeRepository) QueryUnExecutedNotifyTaskList(ctx context.Context) ([]*model.NotifyTaskEntity, error) {
	notifyTaskList, err := r.notifyTaskDAO.QueryUnExecutedNotifyTaskList(ctx)
	if err != nil {
		return nil, err
	}
	if notifyTaskList == nil {
		return []*model.NotifyTaskEntity{}, nil
	}

	notifyTaskEntities := make([]*model.NotifyTaskEntity, 0, len(notifyTaskList))
	for _, notifyTask := range notifyTaskList {
		notifyTaskEntity := &model.NotifyTaskEntity{
			TeamId:        notifyTask.TeamId,
			NotifyUrl:     notifyTask.NotifyUrl,
			NotifyCount:   notifyTask.NotifyCount,
			ParameterJson: notifyTask.ParameterJson,
		}
		notifyTaskEntities = append(notifyTaskEntities, notifyTaskEntity)
	}

	return notifyTaskEntities, nil
}

// QueryUnExecutedNotifyTaskByTeamId 根据团队ID查询未执行的通知任务
func (r *TradeRepository) QueryUnExecutedNotifyTaskByTeamId(ctx context.Context, teamId string) (*model.NotifyTaskEntity, error) {
	notifyTask, err := r.notifyTaskDAO.QueryUnExecutedNotifyTaskByTeamId(ctx, teamId)
	if err != nil {
		return nil, err
	}
	if notifyTask == nil {
		return nil, nil
	}

	return &model.NotifyTaskEntity{
		TeamId:        notifyTask.TeamId,
		NotifyUrl:     notifyTask.NotifyUrl,
		NotifyCount:   notifyTask.NotifyCount,
		ParameterJson: notifyTask.ParameterJson,
	}, nil
}

// UpdateNotifyTaskStatusSuccess 更新通知任务状态为成功
func (r *TradeRepository) UpdateNotifyTaskStatusSuccess(ctx context.Context, teamId string) error {
	return r.notifyTaskDAO.UpdateNotifyTaskStatusSuccess(ctx, teamId)
}

// UpdateNotifyTaskStatusError 更新通知任务状态为错误
func (r *TradeRepository) UpdateNotifyTaskStatusError(ctx context.Context, teamId string) error {
	return r.notifyTaskDAO.UpdateNotifyTaskStatusError(ctx, teamId)
}

// UpdateNotifyTaskStatusRetry 更新通知任务状态为重试
func (r *TradeRepository) UpdateNotifyTaskStatusRetry(ctx context.Context, teamId string) error {
	return r.notifyTaskDAO.UpdateNotifyTaskStatusRetry(ctx, teamId)
}
