package dao

import (
	"context"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/po"
)

// CrowdTagsDetailDAO defines the interface for crowd tags detail persistence
type CrowdTagsDetailDAO interface {
	AddCrowdTagsUserId(ctx context.Context, crowdTagsDetail *po.CrowdTagsDetail) error
}

// MySQLCrowdTagsDetailDAO is a GORM implementation of CrowdTagsDetailDAO
type MySQLCrowdTagsDetailDAO struct {
	data *data.Data
}

// NewMySQLCrowdTagsDetailDAO creates a new MySQL crowd tags detail DAO
func NewMySQLCrowdTagsDetailDAO(data *data.Data) CrowdTagsDetailDAO {
	return &MySQLCrowdTagsDetailDAO{
		data: data,
	}
}

// AddCrowdTagsUserId adds a user ID to crowd tags detail
func (r *MySQLCrowdTagsDetailDAO) AddCrowdTagsUserId(ctx context.Context, crowdTagsDetail *po.CrowdTagsDetail) error {
	return r.data.DB(ctx).WithContext(ctx).Create(crowdTagsDetail).Error
}
