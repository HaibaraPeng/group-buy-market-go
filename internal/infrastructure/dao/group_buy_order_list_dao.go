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
			"status":      1, // 状态1表示已完成
			"update_time": time.Now(),
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
