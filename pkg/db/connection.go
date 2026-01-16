package db

import (
	"devtracker/internal/domain/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init(){
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}

func NewConnection() *gorm.DB{
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "tracker123")
	dbname := getEnv("DB_NAME", "tracker_db")

	 dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
        host, user, password, dbname, port,
    )

    log.Printf("Connecting to database: host=%s port=%s dbname=%s user=%s", host, port, dbname, user)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    fmt.Println("isi dari db:", db)

    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    log.Println("Database connection established successfully")

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Failed to get underlying sql.DB: %v", err)
    }

    if err := sqlDB.Ping(); err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    log.Println("Running database migrations...")
    err = db.AutoMigrate(
        &model.User{},
        &model.Project{},
        &model.Milestone{},
        &model.Log{},
        &model.AIInsight{},
        &model.Report{},
    )
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    log.Println("Database migrations completed successfully")
    return db
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}