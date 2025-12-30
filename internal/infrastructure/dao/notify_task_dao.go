package dao

import (
	"context"
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/po"
	"time"
)

// NotifyTaskDAO defines the interface for notification task persistence
type NotifyTaskDAO interface {
	Insert(ctx context.Context, notifyTask *po.NotifyTask) error
}

// MySQLNotifyTaskDAO is a GORM implementation of NotifyTaskDAO
type MySQLNotifyTaskDAO struct {
	db *gorm.DB
}

// NewMySQLNotifyTaskDAO creates a new MySQL notification task DAO
func NewMySQLNotifyTaskDAO(db *gorm.DB) NotifyTaskDAO {
	return &MySQLNotifyTaskDAO{
		db: db,
	}
}

// Insert inserts a new notification task
func (r *MySQLNotifyTaskDAO) Insert(ctx context.Context, notifyTask *po.NotifyTask) error {
	notifyTask.CreateTime = time.Now()
	notifyTask.UpdateTime = time.Now()
	return r.db.WithContext(ctx).Create(notifyTask).Error
}
