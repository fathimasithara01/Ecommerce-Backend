package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/fathimasithara01/ecommerce/config"
)

var DB *gorm.DB

func ConnectDB(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode= disable", config.DBHost, config.DBUser, config.DBName, config.DBPort, config.DBPaassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database:%w", err)
	}
	DB = db
	return DB, err
}
