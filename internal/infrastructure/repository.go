package infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"group-buy-market-go/internal/domain"
	"time"
)

// MySQLProductRepository is a MySQL implementation of ProductRepository
type MySQLProductRepository struct {
	db *sql.DB
}

// NewMySQLProductRepository creates a new MySQL product repository
func NewMySQLProductRepository(db *sql.DB) *MySQLProductRepository {
	return &MySQLProductRepository{
		db: db,
	}
}

// Save saves a product
func (r *MySQLProductRepository) Save(product *domain.Product) error {
	query := `INSERT INTO products (id, name, price) VALUES (?, ?, ?) 
	          ON DUPLICATE KEY UPDATE name = ?, price = ?`
	_, err := r.db.Exec(query, product.ID, product.Name, product.Price, product.Name, product.Price)
	return err
}

// FindByID finds a product by ID
func (r *MySQLProductRepository) FindByID(id int64) (*domain.Product, error) {
	query := `SELECT id, name, price FROM products WHERE id = ?`
	row := r.db.QueryRow(query, id)

	product := &domain.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return product, nil
}

// FindAll returns all products
func (r *MySQLProductRepository) FindAll() ([]*domain.Product, error) {
	query := `SELECT id, name, price FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		product := &domain.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// MySQLGroupBuyActivityRepository is a MySQL implementation of GroupBuyActivityRepository
type MySQLGroupBuyActivityRepository struct {
	db *sql.DB
}

// NewMySQLGroupBuyActivityRepository creates a new MySQL group buy activity repository
func NewMySQLGroupBuyActivityRepository(db *sql.DB) *MySQLGroupBuyActivityRepository {
	return &MySQLGroupBuyActivityRepository{
		db: db,
	}
}

// Save saves a group buy activity
func (r *MySQLGroupBuyActivityRepository) Save(activity *domain.GroupBuyActivity) error {
	query := `INSERT INTO group_buy_activities 
		(id, activity_id, activity_name, source, channel, goods_id, discount_id, 
		 group_type, take_limit_count, target, valid_time, status, start_time, 
		 end_time, tag_id, tag_scope, create_time, update_time) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) 
		ON DUPLICATE KEY UPDATE 
		activity_name = ?, source = ?, channel = ?, goods_id = ?, discount_id = ?, 
		group_type = ?, take_limit_count = ?, target = ?, valid_time = ?, status = ?, 
		start_time = ?, end_time = ?, tag_id = ?, tag_scope = ?, update_time = ?`

	_, err := r.db.Exec(query,
		activity.ID, activity.ActivityId, activity.ActivityName, activity.Source,
		activity.Channel, activity.GoodsId, activity.DiscountId, activity.GroupType,
		activity.TakeLimitCount, activity.Target, activity.ValidTime, activity.Status,
		activity.StartTime, activity.EndTime, activity.TagId, activity.TagScope,
		activity.CreateTime, activity.UpdateTime,
		// For update
		activity.ActivityName, activity.Source, activity.Channel, activity.GoodsId,
		activity.DiscountId, activity.GroupType, activity.TakeLimitCount, activity.Target,
		activity.ValidTime, activity.Status, activity.StartTime, activity.EndTime,
		activity.TagId, activity.TagScope, activity.UpdateTime)

	return err
}

// FindByID finds a group buy activity by ID
func (r *MySQLGroupBuyActivityRepository) FindByID(id int64) (*domain.GroupBuyActivity, error) {
	query := `SELECT id, activity_id, activity_name, source, channel, goods_id, discount_id, 
	             group_type, take_limit_count, target, valid_time, status, start_time, 
	             end_time, tag_id, tag_scope, create_time, update_time 
			  FROM group_buy_activities WHERE id = ?`
	row := r.db.QueryRow(query, id)

	activity := &domain.GroupBuyActivity{}
	err := row.Scan(
		&activity.ID, &activity.ActivityId, &activity.ActivityName, &activity.Source,
		&activity.Channel, &activity.GoodsId, &activity.DiscountId, &activity.GroupType,
		&activity.TakeLimitCount, &activity.Target, &activity.ValidTime, &activity.Status,
		&activity.StartTime, &activity.EndTime, &activity.TagId, &activity.TagScope,
		&activity.CreateTime, &activity.UpdateTime)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return activity, nil
}

// FindByActivityID finds a group buy activity by activity ID
func (r *MySQLGroupBuyActivityRepository) FindByActivityID(activityID int64) (*domain.GroupBuyActivity, error) {
	query := `SELECT id, activity_id, activity_name, source, channel, goods_id, discount_id, 
	             group_type, take_limit_count, target, valid_time, status, start_time, 
	             end_time, tag_id, tag_scope, create_time, update_time 
			  FROM group_buy_activities WHERE activity_id = ?`
	row := r.db.QueryRow(query, activityID)

	activity := &domain.GroupBuyActivity{}
	err := row.Scan(
		&activity.ID, &activity.ActivityId, &activity.ActivityName, &activity.Source,
		&activity.Channel, &activity.GoodsId, &activity.DiscountId, &activity.GroupType,
		&activity.TakeLimitCount, &activity.Target, &activity.ValidTime, &activity.Status,
		&activity.StartTime, &activity.EndTime, &activity.TagId, &activity.TagScope,
		&activity.CreateTime, &activity.UpdateTime)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return activity, nil
}

// FindAll returns all group buy activities
func (r *MySQLGroupBuyActivityRepository) FindAll() ([]*domain.GroupBuyActivity, error) {
	query := `SELECT id, activity_id, activity_name, source, channel, goods_id, discount_id, 
	             group_type, take_limit_count, target, valid_time, status, start_time, 
	             end_time, tag_id, tag_scope, create_time, update_time 
			  FROM group_buy_activities`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []*domain.GroupBuyActivity
	for rows.Next() {
		activity := &domain.GroupBuyActivity{}
		err := rows.Scan(
			&activity.ID, &activity.ActivityId, &activity.ActivityName, &activity.Source,
			&activity.Channel, &activity.GoodsId, &activity.DiscountId, &activity.GroupType,
			&activity.TakeLimitCount, &activity.Target, &activity.ValidTime, &activity.Status,
			&activity.StartTime, &activity.EndTime, &activity.TagId, &activity.TagScope,
			&activity.CreateTime, &activity.UpdateTime)

		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}

	return activities, nil
}

// UpdateStatus updates the status of a group buy activity
func (r *MySQLGroupBuyActivityRepository) UpdateStatus(id int64, status int) error {
	query := `UPDATE group_buy_activities SET status = ?, update_time = ? WHERE id = ?`
	_, err := r.db.Exec(query, status, time.Now(), id)
	return err
}
