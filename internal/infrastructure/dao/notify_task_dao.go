package dao

import (
	"context"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/po"
	"time"
)

// NotifyTaskDAO defines the interface for notification task persistence
type NotifyTaskDAO interface {
	Insert(ctx context.Context, notifyTask *po.NotifyTask) error
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
