package dao

import (
	"context"
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/po"
)

// CrowdTagsDetailDAO defines the interface for crowd tags detail persistence
type CrowdTagsDetailDAO interface {
	AddCrowdTagsUserId(ctx context.Context, crowdTagsDetail *po.CrowdTagsDetail) error
}

// MySQLCrowdTagsDetailDAO is a GORM implementation of CrowdTagsDetailDAO
type MySQLCrowdTagsDetailDAO struct {
	db *gorm.DB
}

// NewMySQLCrowdTagsDetailDAO creates a new MySQL crowd tags detail DAO
func NewMySQLCrowdTagsDetailDAO(db *gorm.DB) CrowdTagsDetailDAO {
	return &MySQLCrowdTagsDetailDAO{
		db: db,
	}
}

// AddCrowdTagsUserId adds a user ID to crowd tags detail
func (r *MySQLCrowdTagsDetailDAO) AddCrowdTagsUserId(ctx context.Context, crowdTagsDetail *po.CrowdTagsDetail) error {
	return r.db.Create(crowdTagsDetail).Error
}
