package repository

import (
	"FPPBKKGo/internal/domain"

	"gorm.io/gorm"
)

type AlbumRepository struct {
    DB *gorm.DB
}

// Ensure AlbumRepository implements AlbumRepository interface
var _ domain.AlbumRepository = (*AlbumRepository)(nil)

// CreateAlbum creates a new album in the database
func (ar *AlbumRepository) CreateAlbum(album domain.Album) (domain.Album, error) {
    if err := ar.DB.Create(&album).Error; err != nil {
        return domain.Album{}, err
    }
    return album, nil
}

// GetAllAlbums fetches all albums
func (ar *AlbumRepository) GetAllAlbums() ([]domain.Album, error) {
    var albums []domain.Album
    if err := ar.DB.Find(&albums).Error; err != nil {
        return nil, err
    }
    return albums, nil
}

// GetAlbumByID fetches an album by its ID
func (ar *AlbumRepository) GetAlbumByID(id uint) (domain.Album, error) {
    var album domain.Album
    if err := ar.DB.First(&album, id).Error; err != nil {
        return domain.Album{}, err
    }
    return album, nil
}

// UpdateAlbum updates an album in the database
func (ar *AlbumRepository) UpdateAlbum(id uint, album domain.Album) (domain.Album, error) {
    if err := ar.DB.Model(&domain.Album{}).Where("id = ?", id).Updates(album).Error; err != nil {
        return domain.Album{}, err
    }
    return album, nil
}

// DeleteAlbum deletes an album from the database
func (ar *AlbumRepository) DeleteAlbum(id uint) error {
    if err := ar.DB.Delete(&domain.Album{}, id).Error; err != nil {
        return err
    }
    return nil
}
