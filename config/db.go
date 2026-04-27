package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the database connection
func ConnectDB() {
    var err error

    // Build DSN (Data Source Name) from environment variables
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        getEnv("DB_PORT", "5432"),
    )

    // Retry mechanism (useful when DB starts slower than app in Docker)
    maxRetries := 5
    for i := 1; i <= maxRetries; i++ {
        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

        if err == nil {
            log.Println("✅ Database connected successfully")
            break
        }

        log.Printf("❌ Failed to connect to DB (attempt %d/%d): %v", i, maxRetries, err)
        time.Sleep(2 * time.Second)
    }

    if err != nil {
        log.Fatal("🚨 Could not connect to database after retries")
    }

    // Get underlying sql.DB to configure connection pool
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatal("🚨 Failed to get sql.DB from GORM")
    }

    // Set connection pool settings
    sqlDB.SetMaxOpenConns(25)                 // max open connections
    sqlDB.SetMaxIdleConns(10)                 // max idle connections
    sqlDB.SetConnMaxLifetime(5 * time.Minute) // connection lifetime
}

// Helper function to read env with fallback
func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
