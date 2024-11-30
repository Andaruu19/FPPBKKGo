package migrations

import (
	"FPPBKKGo/internal/domain"
	"log"

	"gorm.io/gorm"
)

func MigrateAlbums(db *gorm.DB) {
    err := db.AutoMigrate(&domain.Album{})
    if err != nil {
        log.Fatalf("Failed to migrate genres table: %v", err)
    }
}

func DropAlbumsTable(db *gorm.DB) {
    err := db.Migrator().DropTable(&domain.Album{})
    if err != nil {
        log.Fatalf("Failed to drop genres table: %v", err)
    }
}
