package dao

import (
	"context"
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/po"
)

// CrowdTagsJobDAO defines the interface for crowd tags job persistence
type CrowdTagsJobDAO interface {
	QueryCrowdTagsJob(ctx context.Context, crowdTagsJob *po.CrowdTagsJob) (*po.CrowdTagsJob, error)
}

// MySQLCrowdTagsJobDAO is a GORM implementation of CrowdTagsJobDAO
type MySQLCrowdTagsJobDAO struct {
	db *gorm.DB
}

// NewMySQLCrowdTagsJobDAO creates a new MySQL crowd tags job DAO
func NewMySQLCrowdTagsJobDAO(db *gorm.DB) CrowdTagsJobDAO {
	return &MySQLCrowdTagsJobDAO{
		db: db,
	}
}

// QueryCrowdTagsJob queries crowd tags job by condition
func (r *MySQLCrowdTagsJobDAO) QueryCrowdTagsJob(ctx context.Context, crowdTagsJob *po.CrowdTagsJob) (*po.CrowdTagsJob, error) {
	var result po.CrowdTagsJob
	err := r.db.Where(crowdTagsJob).First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
