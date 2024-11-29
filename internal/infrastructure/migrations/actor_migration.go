package migrations

import (
	"FPPBKKGo/internal/domain"
	"log"

	"gorm.io/gorm"
)

func MigrateActors(db *gorm.DB) {
    err := db.AutoMigrate(&domain.Actor{})
    if err != nil {
        log.Fatalf("Failed to migrate actors table: %v", err)
    }
}

func DropActorsTable(db *gorm.DB) {
    err := db.Migrator().DropTable(&domain.Actor{})
    if err != nil {
        log.Fatalf("Failed to drop actors table: %v", err)
    }
}
