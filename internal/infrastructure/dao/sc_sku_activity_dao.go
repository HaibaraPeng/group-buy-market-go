package dao

import (
	"context"
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/po"
)

// SCSkuActivityDAO defines the interface for sc sku activity persistence
type SCSkuActivityDAO interface {
	QuerySCSkuActivityBySCGoodsId(ctx context.Context, scSkuActivity *po.SCSkuActivity) (*po.SCSkuActivity, error)
}

// MySQLSCSkuActivityDAO is a GORM implementation of SCSkuActivityDAO
type MySQLSCSkuActivityDAO struct {
	data *data.Data
}

// NewMySQLSCSkuActivityDAO creates a new MySQL sc sku activity DAO
func NewMySQLSCSkuActivityDAO(data *data.Data) SCSkuActivityDAO {
	return &MySQLSCSkuActivityDAO{
		data: data,
	}
}

// QuerySCSkuActivityBySCGoodsId queries sc sku activity by source, channel and goods id
func (r *MySQLSCSkuActivityDAO) QuerySCSkuActivityBySCGoodsId(ctx context.Context, scSkuActivity *po.SCSkuActivity) (*po.SCSkuActivity, error) {
	var result po.SCSkuActivity
	err := r.data.DB(ctx).WithContext(ctx).Where("source = ? AND channel = ? AND goods_id = ?", scSkuActivity.Source, scSkuActivity.Channel, scSkuActivity.GoodsId).First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
