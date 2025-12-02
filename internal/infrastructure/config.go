package infrastructure

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
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
	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found: %s", path)
	}

	// Read file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	// Parse YAML
	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return config, nil
}

// InitDB initializes the database connection
func (c *Config) InitDB() (*sql.DB, error) {
	// First connect without specifying a database to create it if needed
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port)

	adminDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open admin database connection: %v", err)
	}
	defer adminDB.Close()

	// Create database if it doesn't exist
	_, err = adminDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", c.Database.DBName))
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %v", err)
	}

	// Now connect to the actual database
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Create tables if they don't exist
	err = c.createTables(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %v", err)
	}

	return db, nil
}

// createTables creates the necessary tables if they don't exist
func (c *Config) createTables(db *sql.DB) error {
	// Create products table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id BIGINT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			price BIGINT NOT NULL
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create products table: %v", err)
	}

	// Create group_buy_activities table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS group_buy_activities (
			id BIGINT PRIMARY KEY,
			activity_id BIGINT NOT NULL,
			activity_name VARCHAR(255) NOT NULL,
			source VARCHAR(50) NOT NULL,
			channel VARCHAR(50) NOT NULL,
			goods_id VARCHAR(50) NOT NULL,
			discount_id VARCHAR(50) NOT NULL,
			group_type INT NOT NULL,
			take_limit_count INT NOT NULL,
			target INT NOT NULL,
			valid_time INT NOT NULL,
			status INT NOT NULL,
			start_time DATETIME,
			end_time DATETIME,
			tag_id VARCHAR(50),
			tag_scope VARCHAR(50),
			create_time DATETIME NOT NULL,
			update_time DATETIME NOT NULL,
			INDEX idx_activity_id (activity_id),
			INDEX idx_status (status)
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create group_buy_activities table: %v", err)
	}

	return nil
}
