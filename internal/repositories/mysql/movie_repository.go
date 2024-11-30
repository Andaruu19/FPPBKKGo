package repository

import (
	"FPPBKKGo/internal/domain"

	"gorm.io/gorm"
)

type MovieRepository struct {
    DB *gorm.DB
}

// FindByID fetches a movie by its ID
func (mr *MovieRepository) GetByID(id uint) (*domain.Movie, error) {
    var movie domain.Movie
    if err := mr.DB.First(&movie, id).Error; err != nil {
        return nil, err
    }
    return &movie, nil
}

// FindAll fetches all movies
func (mr *MovieRepository) GetAll() ([]domain.Movie, error) {
    var movies []domain.Movie
    if err := mr.DB.Find(&movies).Error; err != nil {
        return nil, err
    }
    return movies, nil
}

// GetMoviesByAlbumID fetches all movies associated with a given album ID
func (mr *MovieRepository) GetMoviesByAlbumID(albumID uint) ([]domain.Movie, error) {
	var movies []domain.Movie
	err := mr.DB.Joins("JOIN album_movies ON album_movies.movie_id = movies.id").
		Where("album_movies.album_id = ?", albumID).
		Find(&movies).Error

	if err != nil {
		return nil, err
	}

	return movies, nil
}
