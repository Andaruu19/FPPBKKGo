package migrations

import (
	"gorm.io/gorm"
)

type AlbumMovie struct {
    AlbumID uint `gorm:"not null;index"`
    MovieID uint `gorm:"not null;index"`
}

// MigrateAlbumMovie creates the album_movie table
func MigrateAlbumMovie(db *gorm.DB) error {
    return db.AutoMigrate(&AlbumMovie{})
}
