package domain

// Product represents a product entity
type Product struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"` // price in cents
}

// User represents a user entity
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Order represents an order entity
type Order struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Items  []Item `json:"items"`
	Total  int64  `json:"total"` // total in cents
	Status string `json:"status"`
}

// Item represents an item in an order
type Item struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
	Price     int64 `json:"price"` // price in cents
}
