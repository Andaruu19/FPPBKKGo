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

func (ar *AlbumRepository) CreateAlbum(album domain.Album) (domain.Album, error) {
    if err := ar.DB.Create(&album).Error; err != nil {
        return domain.Album{}, err
    }
    return album, nil
}

func (ar *AlbumRepository) GetAllAlbums() ([]domain.Album, error) {
    var albums []domain.Album
    if err := ar.DB.Find(&albums).Error; err != nil {
        return nil, err
    }
    return albums, nil
}

func (ar *AlbumRepository) GetAlbumByID(id uint) (domain.Album, error) {
    var album domain.Album
    if err := ar.DB.First(&album, id).Error; err != nil {
        return domain.Album{}, err
    }
    return album, nil
}

func (ar *AlbumRepository) UpdateAlbum(id uint, album domain.Album) (domain.Album, error) {
    if err := ar.DB.Model(&domain.Album{}).Where("id = ?", id).Updates(album).Error; err != nil {
        return domain.Album{}, err
    }
    return album, nil
}

func (ar *AlbumRepository) DeleteAlbum(id uint) error {
    if err := ar.DB.Delete(&domain.Album{}, id).Error; err != nil {
        return err
    }
    return nil
}

// AddMovieToAlbum adds a relationship between the album and the movie in the join table
func (ar *AlbumRepository) AddMovieToAlbum(albumID uint, movieID uint) error {
    // Create the relationship in the album_movies join table
    if err := ar.DB.Table("album_movies").Create(&domain.AlbumMovie{
        AlbumID: albumID,
        MovieID: movieID,
    }).Error; err != nil {
        return err
    }
    return nil
}

// RemoveMovieFromAlbum removes the relationship between the album and the movie from the join table
func (ar *AlbumRepository) RemoveMovieFromAlbum(albumID uint, movieID uint) error {
    // Delete the relationship in the album_movies join table
    if err := ar.DB.Table("album_movies").Where("album_id = ? AND movie_id = ?", albumID, movieID).Delete(&domain.AlbumMovie{}).Error; err != nil {
        return err
    }
    return nil
}

