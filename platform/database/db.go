package database

import (
	"FleetManagerAPI/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.DB) error {
    dsn := fmt.Sprintf("host=%s port=%d sslmode=%s user=%s password=%s dbname=%s",
	cfg.Host,
	cfg.Port,
	cfg.SslMode,
	cfg.User,
	cfg.Password,
	cfg.Name,
	)

	var err error
    if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
        return err
    }

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(cfg.MaxConnLifetime)
	fmt.Println("Connection to database established")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}

