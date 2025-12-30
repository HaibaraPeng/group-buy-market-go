package dao

import (
	"context"
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/po"
	"time"
)

// GroupBuyOrderDAO defines the interface for group buy order persistence
type GroupBuyOrderDAO interface {
	Insert(ctx context.Context, groupBuyOrder *po.GroupBuyOrder) error
	UpdateAddLockCount(ctx context.Context, teamId string) (int64, error)
	UpdateSubtractionLockCount(ctx context.Context, teamId string) (int64, error)
	QueryGroupBuyProgress(ctx context.Context, teamId string) (*po.GroupBuyOrder, error)
	QueryGroupBuyTeamByTeamId(ctx context.Context, teamId string) (*po.GroupBuyOrder, error)
	UpdateAddCompleteCount(ctx context.Context, teamId string) (int64, error)
	UpdateOrderStatus2COMPLETE(ctx context.Context, teamId string) (int64, error)
}

// MySQLGroupBuyOrderDAO is a GORM implementation of GroupBuyOrderDAO
type MySQLGroupBuyOrderDAO struct {
	db *gorm.DB
}

// NewMySQLGroupBuyOrderDAO creates a new MySQL group buy order DAO
func NewMySQLGroupBuyOrderDAO(db *gorm.DB) GroupBuyOrderDAO {
	return &MySQLGroupBuyOrderDAO{
		db: db,
	}
}

// Insert inserts a new group buy order
func (r *MySQLGroupBuyOrderDAO) Insert(ctx context.Context, groupBuyOrder *po.GroupBuyOrder) error {
	groupBuyOrder.Status = 0 // 默认状态为0(拼单中)
	groupBuyOrder.CreateTime = time.Now()
	groupBuyOrder.UpdateTime = time.Now()
	return r.db.WithContext(ctx).Create(groupBuyOrder).Error
}

// UpdateAddLockCount increases the lock count for a group buy order
func (r *MySQLGroupBuyOrderDAO) UpdateAddLockCount(ctx context.Context, teamId string) (int64, error) {
	result := r.db.WithContext(ctx).Model(&po.GroupBuyOrder{}).
		Where("team_id = ? AND lock_count < target_count", teamId).
		Updates(map[string]interface{}{
			"lock_count":  gorm.Expr("lock_count + 1"),
			"update_time": time.Now(),
		})
	return result.RowsAffected, result.Error
}

// UpdateSubtractionLockCount decreases the lock count for a group buy order
func (r *MySQLGroupBuyOrderDAO) UpdateSubtractionLockCount(ctx context.Context, teamId string) (int64, error) {
	result := r.db.WithContext(ctx).Model(&po.GroupBuyOrder{}).
		Where("team_id = ? AND lock_count > 0", teamId).
		Updates(map[string]interface{}{
			"lock_count":  gorm.Expr("lock_count - 1"),
			"update_time": time.Now(),
		})
	return result.RowsAffected, result.Error
}

// QueryGroupBuyProgress queries the progress of a group buy order
func (r *MySQLGroupBuyOrderDAO) QueryGroupBuyProgress(ctx context.Context, teamId string) (*po.GroupBuyOrder, error) {
	var groupBuyOrder po.GroupBuyOrder
	err := r.db.WithContext(ctx).Select("target_count", "complete_count", "lock_count").
		Where("team_id = ?", teamId).
		First(&groupBuyOrder).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &groupBuyOrder, nil
}

// QueryGroupBuyTeamByTeamId queries the team information by teamId
func (r *MySQLGroupBuyOrderDAO) QueryGroupBuyTeamByTeamId(ctx context.Context, teamId string) (*po.GroupBuyOrder, error) {
	var groupBuyOrder po.GroupBuyOrder
	err := r.db.WithContext(ctx).Select("team_id, activity_id, target_count, complete_count, lock_count, status").
		Where("team_id = ?", teamId).
		First(&groupBuyOrder).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &groupBuyOrder, nil
}

// UpdateAddCompleteCount increases the complete count for a group buy order
func (r *MySQLGroupBuyOrderDAO) UpdateAddCompleteCount(ctx context.Context, teamId string) (int64, error) {
	result := r.db.WithContext(ctx).Model(&po.GroupBuyOrder{}).
		Where("team_id = ? AND complete_count < target_count", teamId).
		Updates(map[string]interface{}{
			"complete_count": gorm.Expr("complete_count + 1"),
			"update_time":    time.Now(),
		})
	return result.RowsAffected, result.Error
}

// UpdateOrderStatus2COMPLETE updates the order status to COMPLETE
func (r *MySQLGroupBuyOrderDAO) UpdateOrderStatus2COMPLETE(ctx context.Context, teamId string) (int64, error) {
	result := r.db.WithContext(ctx).Model(&po.GroupBuyOrder{}).
		Where("team_id = ? AND status = 0", teamId).
		Updates(map[string]interface{}{
			"status":      1, // 状态1表示已完成
			"update_time": time.Now(),
		})
	return result.RowsAffected, result.Error
}
