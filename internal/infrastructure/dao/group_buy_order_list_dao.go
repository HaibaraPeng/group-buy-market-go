package dao

import (
	"context"
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/po"
	"time"
)

// GroupBuyOrderListDAO defines the interface for group buy order list persistence
type GroupBuyOrderListDAO interface {
	Insert(ctx context.Context, groupBuyOrderList *po.GroupBuyOrderList) error
	QueryGroupBuyOrderRecordByOutTradeNo(ctx context.Context, req *po.GroupBuyOrderList) (*po.GroupBuyOrderList, error)
	QueryOrderCountByActivityId(ctx context.Context, req *po.GroupBuyOrderList) (int, error) // 添加查询活动订单数的方法
}

// MySQLGroupBuyOrderListDAO is a GORM implementation of GroupBuyOrderListDAO
type MySQLGroupBuyOrderListDAO struct {
	db *gorm.DB
}

// NewMySQLGroupBuyOrderListDAO creates a new MySQL group buy order list DAO
func NewMySQLGroupBuyOrderListDAO(db *gorm.DB) GroupBuyOrderListDAO {
	return &MySQLGroupBuyOrderListDAO{
		db: db,
	}
}

// Insert inserts a new group buy order list record
func (r *MySQLGroupBuyOrderListDAO) Insert(ctx context.Context, groupBuyOrderList *po.GroupBuyOrderList) error {
	groupBuyOrderList.CreateTime = time.Now()
	groupBuyOrderList.UpdateTime = time.Now()
	return r.db.WithContext(ctx).Create(groupBuyOrderList).Error
}

// QueryGroupBuyOrderRecordByOutTradeNo queries group buy order record by out trade number
func (r *MySQLGroupBuyOrderListDAO) QueryGroupBuyOrderRecordByOutTradeNo(ctx context.Context, req *po.GroupBuyOrderList) (*po.GroupBuyOrderList, error) {
	var groupBuyOrderList po.GroupBuyOrderList
	err := r.db.WithContext(ctx).Select("user_id", "team_id", "order_id", "activity_id", "start_time",
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
	err := r.db.WithContext(ctx).Model(&po.GroupBuyOrderList{}).
		Where("activity_id = ? AND user_id = ?", req.ActivityId, req.UserId).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
