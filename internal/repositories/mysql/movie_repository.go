package repository

import (
	"FPPBKKGo/internal/domain"

	"gorm.io/gorm"
)

type MovieRepository struct {
    DB *gorm.DB
}

// FindByID fetches a movie by its ID
func (mr *MovieRepository) FindByID(id uint) (*domain.Movie, error) {
    var movie domain.Movie
    if err := mr.DB.First(&movie, id).Error; err != nil {
        return nil, err
    }
    return &movie, nil
}

// FindAll fetches all movies
func (mr *MovieRepository) FindAll() ([]domain.Movie, error) {
    var movies []domain.Movie
    if err := mr.DB.Find(&movies).Error; err != nil {
        return nil, err
    }
    return movies, nil
}
