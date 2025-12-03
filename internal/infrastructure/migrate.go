package infrastructure

import (
	"group-buy-market-go/internal/domain"
	"gorm.io/gorm"
)

// Migrate runs database migrations for all models
func Migrate(db *gorm.DB) error {
	// Run auto-migration for all models
	err := db.AutoMigrate(
		&domain.Product{},
		&domain.User{},
		&domain.Order{},
		&domain.Item{},
		&domain.GroupBuyActivity{},
	)
	return err
}