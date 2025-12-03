package infrastructure

import (
	"gorm.io/gorm"
	"group-buy-market-go/internal/domain"
	"group-buy-market-go/internal/infrastructure/po"
)

// Migrate runs database migrations for all models
func Migrate(db *gorm.DB) error {
	// Run auto-migration for all models
	err := db.AutoMigrate(
		&domain.Product{},
		&domain.User{},
		&domain.Order{},
		&domain.Item{},
		&po.GroupBuyActivity{},
	)
	return err
}
