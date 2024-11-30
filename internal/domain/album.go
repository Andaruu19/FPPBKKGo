package domain

import "gorm.io/gorm"

// Album represents the album entity
type Album struct {
    gorm.Model
    Slug      string   `gorm:"type:varchar(255);uniqueIndex;not null"`
    Name      string   `gorm:"type:varchar(255);not null"`
    Deskripsi string   `gorm:"type:text;not null"`
    Movies    []*Movie `gorm:"many2many:album_movies;"` // For many-to-many relationship
}

type AlbumRepository interface {
    CreateAlbum(album Album) (Album, error)
    GetAlbumByID(id uint) (Album, error)
    GetAllAlbums() ([]Album, error)
    UpdateAlbum(id uint, album Album) (Album, error)
    DeleteAlbum(id uint) error
}
