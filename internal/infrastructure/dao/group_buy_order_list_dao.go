package dao

import (
	"context"
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/po"
	"time"
)

// GroupBuyOrderListDAO defines the interface for group buy order list persistence
type GroupBuyOrderListDAO interface {
	Insert(ctx context.Context, groupBuyOrderList *po.GroupBuyOrderList) error
	QueryGroupBuyOrderRecordByOutTradeNo(ctx context.Context, req *po.GroupBuyOrderList) (*po.GroupBuyOrderList, error)
	QueryOrderCountByActivityId(ctx context.Context, req *po.GroupBuyOrderList) (int, error)
	UpdateOrderStatus2COMPLETE(ctx context.Context, req *po.GroupBuyOrderList) (int64, error)
	QueryGroupBuyCompleteOrderOutTradeNoListByTeamId(ctx context.Context, teamId string) ([]string, error)
	QueryInProgressUserGroupBuyOrderDetailListByUserId(ctx context.Context, req *po.GroupBuyOrderList) ([]*po.GroupBuyOrderList, error)
	QueryInProgressUserGroupBuyOrderDetailListByRandom(ctx context.Context, req *po.GroupBuyOrderList) ([]*po.GroupBuyOrderList, error)
	QueryInProgressUserGroupBuyOrderDetailListByActivityId(ctx context.Context, activityId int64) ([]*po.GroupBuyOrderList, error)
}

// MySQLGroupBuyOrderListDAO is a GORM implementation of GroupBuyOrderListDAO
type MySQLGroupBuyOrderListDAO struct {
	data *data.Data
}

// NewMySQLGroupBuyOrderListDAO creates a new MySQL group buy order list DAO
func NewMySQLGroupBuyOrderListDAO(data *data.Data) GroupBuyOrderListDAO {
	return &MySQLGroupBuyOrderListDAO{
		data: data,
	}
}

// Insert inserts a new group buy order list record
func (r *MySQLGroupBuyOrderListDAO) Insert(ctx context.Context, groupBuyOrderList *po.GroupBuyOrderList) error {
	groupBuyOrderList.CreateTime = time.Now()
	groupBuyOrderList.UpdateTime = time.Now()
	return r.data.DB(ctx).WithContext(ctx).Create(groupBuyOrderList).Error
}

// QueryGroupBuyOrderRecordByOutTradeNo queries group buy order record by out trade number
func (r *MySQLGroupBuyOrderListDAO) QueryGroupBuyOrderRecordByOutTradeNo(ctx context.Context, req *po.GroupBuyOrderList) (*po.GroupBuyOrderList, error) {
	var groupBuyOrderList po.GroupBuyOrderList
	err := r.data.DB(ctx).WithContext(ctx).Select("user_id", "team_id", "order_id", "activity_id", "start_time",
		"end_time", "goods_id", "source", "channel", "original_price", "deduction_price", "status").
		Where("out_trade_no = ? AND user_id = ? AND status = ?", req.OutTradeNo, req.UserId, 0).
		First(&groupBuyOrderList).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &groupBuyOrderList, nil
}

// QueryOrderCountByActivityId queries order count by activity ID and user ID
func (r *MySQLGroupBuyOrderListDAO) QueryOrderCountByActivityId(ctx context.Context, req *po.GroupBuyOrderList) (int, error) {
	var count int64
	err := r.data.DB(ctx).WithContext(ctx).Model(&po.GroupBuyOrderList{}).
		Where("activity_id = ? AND user_id = ?", req.ActivityId, req.UserId).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

// UpdateOrderStatus2COMPLETE updates the order status to COMPLETE
func (r *MySQLGroupBuyOrderListDAO) UpdateOrderStatus2COMPLETE(ctx context.Context, req *po.GroupBuyOrderList) (int64, error) {
	result := r.data.DB(ctx).WithContext(ctx).Model(&po.GroupBuyOrderList{}).
		Where("out_trade_no = ? AND user_id = ?", req.OutTradeNo, req.UserId).
		Updates(map[string]interface{}{
			"status":         1, // 状态1表示已完成
			"out_trade_time": req.OutTradeTime,
			"update_time":    time.Now(),
		})
	return result.RowsAffected, result.Error
}

// QueryGroupBuyCompleteOrderOutTradeNoListByTeamId queries the list of completed order transaction numbers by team ID
func (r *MySQLGroupBuyOrderListDAO) QueryGroupBuyCompleteOrderOutTradeNoListByTeamId(ctx context.Context, teamId string) ([]string, error) {
	var outTradeNos []string
	err := r.data.DB(ctx).WithContext(ctx).Model(&po.GroupBuyOrderList{}).
		Select("out_trade_no").
		Where("team_id = ? AND status = 1", teamId).
		Scan(&outTradeNos).Error
	if err != nil {
		return nil, err
	}
	return outTradeNos, nil
}

// QueryInProgressUserGroupBuyOrderDetailListByUserId queries the list of in-progress user group buy order details by user ID
func (r *MySQLGroupBuyOrderListDAO) QueryInProgressUserGroupBuyOrderDetailListByUserId(ctx context.Context, req *po.GroupBuyOrderList) ([]*po.GroupBuyOrderList, error) {
	var groupBuyOrderList []*po.GroupBuyOrderList
	err := r.data.DB(ctx).WithContext(ctx).Select("user_id", "team_id", "out_trade_no").
		Where("activity_id = ? AND user_id = ? AND status IN (0, 1) AND end_time > ?", req.ActivityId, req.UserId, time.Now()).
		Order("id DESC").
		Limit(int(req.Count)).
		Find(&groupBuyOrderList).Error
	if err != nil {
		return nil, err
	}
	return groupBuyOrderList, nil
}

// QueryInProgressUserGroupBuyOrderDetailListByRandom queries the list of in-progress user group buy order details randomly
func (r *MySQLGroupBuyOrderListDAO) QueryInProgressUserGroupBuyOrderDetailListByRandom(ctx context.Context, req *po.GroupBuyOrderList) ([]*po.GroupBuyOrderList, error) {
	var groupBuyOrderList []*po.GroupBuyOrderList
	err := r.data.DB(ctx).WithContext(ctx).Select("user_id", "team_id", "out_trade_no").
		Where("activity_id = ? AND team_id IN (SELECT team_id FROM group_buy_order WHERE activity_id = ? AND status = 0) AND user_id != ? AND status IN (0, 1) AND end_time > ?",
			req.ActivityId, req.ActivityId, req.UserId, time.Now()).
		Order("id DESC").
		Limit(int(req.Count)).
		Find(&groupBuyOrderList).Error
	if err != nil {
		return nil, err
	}
	return groupBuyOrderList, nil
}

// QueryInProgressUserGroupBuyOrderDetailListByActivityId queries the list of in-progress user group buy order details by activity ID
func (r *MySQLGroupBuyOrderListDAO) QueryInProgressUserGroupBuyOrderDetailListByActivityId(ctx context.Context, activityId int64) ([]*po.GroupBuyOrderList, error) {
	var groupBuyOrderList []*po.GroupBuyOrderList
	err := r.data.DB(ctx).WithContext(ctx).Select("user_id", "team_id", "out_trade_no").
		Where("activity_id = ? AND status IN (0, 1)", activityId).
		Group("team_id").
		Find(&groupBuyOrderList).Error
	if err != nil {
		return nil, err
	}
	return groupBuyOrderList, nil
}
