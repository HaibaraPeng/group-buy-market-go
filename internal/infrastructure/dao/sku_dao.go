package dao

import (
	"context"
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/po"
)

// SkuDAO defines the interface for sku persistence
type SkuDAO interface {
	FindByGoodsId(ctx context.Context, goodsId string) (*po.Sku, error)
}

// MySQLSkuDAO is a GORM implementation of SkuDAO
type MySQLSkuDAO struct {
	data *data.Data
}

// NewMySQLSkuDAO creates a new MySQL sku DAO
func NewMySQLSkuDAO(data *data.Data) SkuDAO {
	return &MySQLSkuDAO{
		data: data,
	}
}

// FindByGoodsId finds a sku by goods id
func (r *MySQLSkuDAO) FindByGoodsId(ctx context.Context, goodsId string) (*po.Sku, error) {
	var sku po.Sku
	err := r.data.DB(ctx).WithContext(ctx).Where("goods_id = ?", goodsId).First(&sku).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &sku, nil
}
