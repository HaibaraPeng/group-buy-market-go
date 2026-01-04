package dao

import (
	"context"
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/po"
	"time"
)

// NotifyTaskDAO defines the interface for notification task persistence
type NotifyTaskDAO interface {
	Insert(ctx context.Context, notifyTask *po.NotifyTask) error
	QueryUnExecutedNotifyTaskList(ctx context.Context) ([]*po.NotifyTask, error)
	QueryUnExecutedNotifyTaskByTeamId(ctx context.Context, teamId string) (*po.NotifyTask, error)
	UpdateNotifyTaskStatusSuccess(ctx context.Context, teamId string) error
	UpdateNotifyTaskStatusError(ctx context.Context, teamId string) error
	UpdateNotifyTaskStatusRetry(ctx context.Context, teamId string) error
}

// MySQLNotifyTaskDAO is a GORM implementation of NotifyTaskDAO
type MySQLNotifyTaskDAO struct {
	data *data.Data
}

// NewMySQLNotifyTaskDAO creates a new MySQL notification task DAO
func NewMySQLNotifyTaskDAO(data *data.Data) NotifyTaskDAO {
	return &MySQLNotifyTaskDAO{
		data: data,
	}
}

// Insert inserts a new notification task
func (r *MySQLNotifyTaskDAO) Insert(ctx context.Context, notifyTask *po.NotifyTask) error {
	notifyTask.CreateTime = time.Now()
	notifyTask.UpdateTime = time.Now()
	return r.data.DB(ctx).WithContext(ctx).Create(notifyTask).Error
}

// QueryUnExecutedNotifyTaskList queries unexecuted notification tasks (status 0 or 2)
func (r *MySQLNotifyTaskDAO) QueryUnExecutedNotifyTaskList(ctx context.Context) ([]*po.NotifyTask, error) {
	var notifyTasks []*po.NotifyTask
	err := r.data.DB(ctx).WithContext(ctx).
		Select("activity_id, team_id, notify_url, notify_count, notify_status, parameter_json").
		Where("notify_status IN ?", []int{0, 2}).
		Limit(50).
		Find(&notifyTasks).Error
	return notifyTasks, err
}

// QueryUnExecutedNotifyTaskByTeamId queries unexecuted notification task by team ID
func (r *MySQLNotifyTaskDAO) QueryUnExecutedNotifyTaskByTeamId(ctx context.Context, teamId string) (*po.NotifyTask, error) {
	var notifyTask *po.NotifyTask
	err := r.data.DB(ctx).WithContext(ctx).
		Select("activity_id, team_id, notify_url, notify_count, notify_status, parameter_json").
		Where("team_id = ? AND notify_status IN ?", teamId, []int{0, 2}).
		First(&notifyTask).Error
	return notifyTask, err
}

// UpdateNotifyTaskStatusSuccess updates notification task status to success
func (r *MySQLNotifyTaskDAO) UpdateNotifyTaskStatusSuccess(ctx context.Context, teamId string) error {
	return r.data.DB(ctx).WithContext(ctx).
		Model(&po.NotifyTask{}).
		Where("team_id = ?", teamId).
		Updates(map[string]interface{}{
			"notify_count":  gorm.Expr("notify_count + 1"),
			"notify_status": 1,
			"update_time":   time.Now(),
		}).Error
}

// UpdateNotifyTaskStatusError updates notification task status to error
func (r *MySQLNotifyTaskDAO) UpdateNotifyTaskStatusError(ctx context.Context, teamId string) error {
	return r.data.DB(ctx).WithContext(ctx).
		Model(&po.NotifyTask{}).
		Where("team_id = ?", teamId).
		Updates(map[string]interface{}{
			"notify_count":  gorm.Expr("notify_count + 1"),
			"notify_status": 3,
			"update_time":   time.Now(),
		}).Error
}

// UpdateNotifyTaskStatusRetry updates notification task status to retry
func (r *MySQLNotifyTaskDAO) UpdateNotifyTaskStatusRetry(ctx context.Context, teamId string) error {
	return r.data.DB(ctx).WithContext(ctx).
		Model(&po.NotifyTask{}).
		Where("team_id = ?", teamId).
		Updates(map[string]interface{}{
			"notify_count":  gorm.Expr("notify_count + 1"),
			"notify_status": 2,
			"update_time":   time.Now(),
		}).Error
}
