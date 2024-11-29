package main

import (
	"FPPBKKGo/internal/infrastructure"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    // Connect to MySQL server without specifying a database
    dsn := "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Create the database
    dbName := "FpGo"
    err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error
    if err != nil {
        log.Fatalf("Failed to create database: %v", err)
    }

    // Connect to the newly created database
    dsn = "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Run database migrations
    infrastructure.MigrateDB(db)
}