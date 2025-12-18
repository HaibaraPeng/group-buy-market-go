package dao

import (
	"gorm.io/gorm"
	"group-buy-market-go/internal/infrastructure/po"
	"time"
)

// GroupBuyDiscountDAO defines the interface for group buy discount persistence
type GroupBuyDiscountDAO interface {
	Save(discount *po.GroupBuyDiscount) error
	FindByID(id int64) (*po.GroupBuyDiscount, error)
	FindByDiscountID(discountID int) (*po.GroupBuyDiscount, error)
	QueryGroupBuyDiscountList() ([]*po.GroupBuyDiscount, error)
	Update(discount *po.GroupBuyDiscount) error
}

// MySQLGroupBuyDiscountDAO is a GORM implementation of GroupBuyDiscountDAO
type MySQLGroupBuyDiscountDAO struct {
	db *gorm.DB
}

// NewMySQLGroupBuyDiscountDAO creates a new MySQL group buy discount DAO
func NewMySQLGroupBuyDiscountDAO(db *gorm.DB) GroupBuyDiscountDAO {
	return &MySQLGroupBuyDiscountDAO{
		db: db,
	}
}

// Save saves a group buy discount
func (r *MySQLGroupBuyDiscountDAO) Save(discount *po.GroupBuyDiscount) error {
	return r.db.Save(discount).Error
}

// FindByID finds a group buy discount by ID
func (r *MySQLGroupBuyDiscountDAO) FindByID(id int64) (*po.GroupBuyDiscount, error) {
	var discount po.GroupBuyDiscount
	err := r.db.First(&discount, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &discount, nil
}

// FindByDiscountID finds a group buy discount by discount ID
func (r *MySQLGroupBuyDiscountDAO) FindByDiscountID(discountID int) (*po.GroupBuyDiscount, error) {
	var discount po.GroupBuyDiscount
	err := r.db.Where("discount_id = ?", discountID).First(&discount).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &discount, nil
}

// FindAll returns all group buy discounts
func (r *MySQLGroupBuyDiscountDAO) QueryGroupBuyDiscountList() ([]*po.GroupBuyDiscount, error) {
	var discounts []*po.GroupBuyDiscount
	err := r.db.Find(&discounts).Error
	return discounts, err
}

// Update updates a group buy discount
func (r *MySQLGroupBuyDiscountDAO) Update(discount *po.GroupBuyDiscount) error {
	return r.db.Model(&po.GroupBuyDiscount{}).Where("id = ?", discount.ID).Updates(discount).Update("update_time", time.Now()).Error
}
