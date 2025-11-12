package config

import (
	"fmt"
	"log"
	"pustaka-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(config *Config) *gorm.DB {
	// Build DSN from config
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
		config.DBSSLMode,
		config.DBTimezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	log.Println("Database connection succeed")

	// Auto migration
	db.AutoMigrate(&models.Book{}, &models.User{}, &models.Authentication{})

	return db
}
