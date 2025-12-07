package dao

import (
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/po"
)

// SkuDAO defines the interface for sku persistence
type SkuDAO interface {
	FindByGoodsId(goodsId string) (*po.Sku, error)
}

// MySQLSkuDAO is a GORM implementation of SkuDAO
type MySQLSkuDAO struct {
	db *gorm.DB
}

// NewMySQLSkuDAO creates a new MySQL sku DAO
func NewMySQLSkuDAO(db *gorm.DB) *MySQLSkuDAO {
	return &MySQLSkuDAO{
		db: db,
	}
}

// FindByGoodsId finds a sku by goods id
func (r *MySQLSkuDAO) FindByGoodsId(goodsId string) (*po.Sku, error) {
	var sku po.Sku
	err := r.db.Where("goods_id = ?", goodsId).First(&sku).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &sku, nil
}
