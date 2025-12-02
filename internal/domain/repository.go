package domain

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
