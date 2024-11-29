package migrations

import (
	"FPPBKKGo/internal/domain"
	"log"

	"gorm.io/gorm"
)

func MigrateMovies(db *gorm.DB) {
    err := db.AutoMigrate(&domain.Movie{})
    if err != nil {
        log.Fatalf("Failed to migrate movies table: %v", err)
    }
}

func DropMoviesTable(db *gorm.DB) {
    err := db.Migrator().DropTable(&domain.Movie{})
    if err != nil {
        log.Fatalf("Failed to drop movies table: %v", err)
    }
}
