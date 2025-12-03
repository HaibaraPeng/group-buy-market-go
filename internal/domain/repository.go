package domain

import "group-buy-market-go/internal/infrastructure/po"

// ProductRepository defines the interface for product persistence
type ProductRepository interface {
	Save(product *Product) error
	FindByID(id int64) (*Product, error)
	FindAll() ([]*Product, error)
}

// UserRepository defines the interface for user persistence
type UserRepository interface {
	Save(user *User) error
	FindByID(id int64) (*User, error)
	FindByUsername(username string) (*User, error)
}

// OrderRepository defines the interface for order persistence
type OrderRepository interface {
	Save(order *Order) error
	FindByID(id int64) (*Order, error)
	FindByUserID(userID int64) ([]*Order, error)
}

// GroupBuyActivityRepository defines the interface for group buy activity persistence
type GroupBuyActivityRepository interface {
	Save(activity *po.GroupBuyActivity) error
	FindByID(id int64) (*po.GroupBuyActivity, error)
	FindByActivityID(activityID int64) (*po.GroupBuyActivity, error)
	FindAll() ([]*po.GroupBuyActivity, error)
	UpdateStatus(id int64, status int) error
}
