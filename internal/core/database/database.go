package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	// Database global variable database
	Database = &gorm.DB{}
)

// Session session
type Session struct {
	Database *gorm.DB
}

// Configuration config mysql
type Configuration struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Debug    bool
}

// New new database connection
func New(c *Configuration) (*Session, error) {
	dns := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s  sslmode=disable password=%s",
		c.Host,
		c.Port,
		c.Username,
		c.Name,
		c.Password,
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if c.Debug {
		db = db.Debug()
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return &Session{
		Database: db,
	}, nil
}

// Close close session
func (s *Session) Close() {
}
