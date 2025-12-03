package domain

// Product represents a product entity
type Product struct {
	ID    int64  `json:"id" gorm:"primaryKey;column:id"`
	Name  string `json:"name" gorm:"column:name"`
	Price int64  `json:"price" gorm:"column:price"` // price in cents
}

// TableName specifies the table name for Product
func (Product) TableName() string {
	return "products"
}

// User represents a user entity
type User struct {
	ID       int64  `json:"id" gorm:"primaryKey;column:id"`
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
}

// TableName specifies the table name for User
func (User) TableName() string {
	return "users"
}

// Order represents an order entity
type Order struct {
	ID     int64  `json:"id" gorm:"primaryKey;column:id"`
	UserID int64  `json:"user_id" gorm:"column:user_id"`
	Items  []Item `json:"items" gorm:"foreignKey:OrderID"`
	Total  int64  `json:"total" gorm:"column:total"` // total in cents
	Status string `json:"status" gorm:"column:status"`
}

// TableName specifies the table name for Order
func (Order) TableName() string {
	return "orders"
}

// Item represents an item in an order
type Item struct {
	ID        int64 `json:"id" gorm:"primaryKey;column:id"`
	OrderID   int64 `json:"order_id" gorm:"column:order_id"`
	ProductID int64 `json:"product_id" gorm:"column:product_id"`
	Quantity  int   `json:"quantity" gorm:"column:quantity"`
	Price     int64 `json:"price" gorm:"column:price"` // price in cents
}

// TableName specifies the table name for Item
func (Item) TableName() string {
	return "order_items"
}
