package dao

import (
	"context"
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/po"
	"time"
)

// CrowdTagsDAO defines the interface for crowd tags persistence
type CrowdTagsDAO interface {
	UpdateCrowdTagsStatistics(ctx context.Context, crowdTags *po.CrowdTags) error
}

// MySQLCrowdTagsDAO is a GORM implementation of CrowdTagsDAO
type MySQLCrowdTagsDAO struct {
	db *gorm.DB
}

// NewMySQLCrowdTagsDAO creates a new MySQL crowd tags DAO
func NewMySQLCrowdTagsDAO(db *gorm.DB) CrowdTagsDAO {
	return &MySQLCrowdTagsDAO{
		db: db,
	}
}

// UpdateCrowdTagsStatistics updates crowd tags statistics
func (r *MySQLCrowdTagsDAO) UpdateCrowdTagsStatistics(ctx context.Context, crowdTags *po.CrowdTags) error {
	return r.db.Model(&po.CrowdTags{}).Where("tag_id = ?", crowdTags.TagId).Updates(map[string]interface{}{
		"statistics":  crowdTags.Statistics,
		"update_time": time.Now(),
	}).Error
}
