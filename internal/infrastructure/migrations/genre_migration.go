package migrations

import (
	"FPPBKKGo/internal/domain"
	"log"

	"gorm.io/gorm"
)

func MigrateGenres(db *gorm.DB) {
    err := db.AutoMigrate(&domain.Genre{})
    if err != nil {
        log.Fatalf("Failed to migrate genres table: %v", err)
    }
}

func DropGenresTable(db *gorm.DB) {
    err := db.Migrator().DropTable(&domain.Genre{})
    if err != nil {
        log.Fatalf("Failed to drop genres table: %v", err)
    }
}