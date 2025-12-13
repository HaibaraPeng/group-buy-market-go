package infrastructure

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config holds the application configuration
type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`

	Redis struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"redis"`

	Logging struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
	} `yaml:"logging"`
}

// LoadConfig loads configuration from file
func LoadConfig(path string) (*Config, error) {
	// Create a config instance with file source
	c := config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)

	// Load the config
	if err := c.Load(); err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	// Create config struct and unmarshal
	var cfg Config
	if err := c.Scan(&cfg); err != nil {
		return nil, fmt.Errorf("failed to scan config: %v", err)
	}

	return &cfg, nil
}

// InitDB initializes the database connection with GORM
func (c *Config) InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}
