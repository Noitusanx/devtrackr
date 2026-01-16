package main

import (
	"devtracker/pkg/db"
	"log"
)


func main() {
    log.Println("Testing database connection...")
    
    database := db.NewConnection()
    
    sqlDB, err := database.DB()
    if err != nil {
        log.Fatal("Failed to get underlying sql.DB:", err)
    }
    defer sqlDB.Close()

    if err := sqlDB.Ping(); err != nil {
        log.Fatal("Failed to ping database:", err)
    }

    log.Println("✅ Database connection test successful!")
    log.Println("✅ Auto-migration completed!")
    log.Println("✅ Ready to start development!")
}