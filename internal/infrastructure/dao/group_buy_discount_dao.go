package dao

import (
	"context"
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/po"
)

// GroupBuyDiscountDAO defines the interface for group buy discount persistence
type GroupBuyDiscountDAO interface {
	FindByDiscountID(ctx context.Context, discountID string) (*po.GroupBuyDiscount, error)
	QueryGroupBuyActivityDiscountByDiscountId(ctx context.Context, discountId string) (*po.GroupBuyDiscount, error)
}

// MySQLGroupBuyDiscountDAO is a GORM implementation of GroupBuyDiscountDAO
type MySQLGroupBuyDiscountDAO struct {
	data *data.Data
}

// NewMySQLGroupBuyDiscountDAO creates a new MySQL group buy discount DAO
func NewMySQLGroupBuyDiscountDAO(data *data.Data) GroupBuyDiscountDAO {
	return &MySQLGroupBuyDiscountDAO{
		data: data,
	}
}

// FindByDiscountID finds a group buy discount by discount ID
func (r *MySQLGroupBuyDiscountDAO) FindByDiscountID(ctx context.Context, discountID string) (*po.GroupBuyDiscount, error) {
	var discount po.GroupBuyDiscount
	err := r.data.DB(ctx).WithContext(ctx).Where("discount_id = ?", discountID).First(&discount).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &discount, nil
}

// QueryGroupBuyActivityDiscountByDiscountId finds a group buy discount by discount ID
func (r *MySQLGroupBuyDiscountDAO) QueryGroupBuyActivityDiscountByDiscountId(ctx context.Context, discountId string) (*po.GroupBuyDiscount, error) {
	var discount po.GroupBuyDiscount
	err := r.data.DB(ctx).WithContext(ctx).Where("discount_id = ?", discountId).First(&discount).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &discount, nil
}
