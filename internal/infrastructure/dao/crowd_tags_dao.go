package dao

import (
	"context"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/po"
	"time"
)

// CrowdTagsDAO defines the interface for crowd tags persistence
type CrowdTagsDAO interface {
	UpdateCrowdTagsStatistics(ctx context.Context, crowdTags *po.CrowdTags) error
}

// MySQLCrowdTagsDAO is a GORM implementation of CrowdTagsDAO
type MySQLCrowdTagsDAO struct {
	data *data.Data
}

// NewMySQLCrowdTagsDAO creates a new MySQL crowd tags DAO
func NewMySQLCrowdTagsDAO(data *data.Data) CrowdTagsDAO {
	return &MySQLCrowdTagsDAO{
		data: data,
	}
}

// UpdateCrowdTagsStatistics updates crowd tags statistics
func (r *MySQLCrowdTagsDAO) UpdateCrowdTagsStatistics(ctx context.Context, crowdTags *po.CrowdTags) error {
	return r.data.DB(ctx).WithContext(ctx).Model(&po.CrowdTags{}).Where("tag_id = ?", crowdTags.TagId).Updates(map[string]interface{}{
		"statistics":  crowdTags.Statistics,
		"update_time": time.Now(),
	}).Error
}
