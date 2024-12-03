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
	if err := mr.DB.Preload("Genre").Preload("Actor").First(&movie, id).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

// FindAll fetches all movies
func (mr *MovieRepository) GetAll() ([]domain.Movie, error) {
	var movies []domain.Movie
	if err := mr.DB.Preload("Genre").Preload("Actor").Find(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}

// GetMoviesByAlbumID fetches all movies associated with a given album ID
func (mr *MovieRepository) GetMoviesByAlbumID(albumID uint) ([]domain.Movie, error) {
	var movies []domain.Movie
	err := mr.DB.
		Preload("Genre").
		Preload("Actor").
		Joins("JOIN album_movies ON album_movies.movie_id = movies.id").
		Where("album_movies.album_id = ?", albumID).
		Find(&movies).Error

	if err != nil {
		return nil, err
	}

	return movies, nil
}

// Add to movie_repository.go
func (r *MovieRepository) GetBySlug(slug string) (*domain.Movie, error) {
	var movie domain.Movie
	result := r.DB.Preload("Genre").Where("slug = ?", slug).First(&movie)
	if result.Error != nil {
		return nil, result.Error
	}
	return &movie, nil
}

func (mr *MovieRepository) GetByGenre(slug string) ([]domain.Movie, error) {
	var movies []domain.Movie
	err := mr.DB.
		Preload("Genre").
		Preload("Actor").
		Joins("JOIN genres ON genres.id = movies.genre_id").
		Where("genres.slug = ?", slug).
		Find(&movies).Error

	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (mr *MovieRepository) GetByActor(id uint) ([]domain.Movie, error) {
	var movies []domain.Movie
	err := mr.DB.
		Preload("Genre").
		Preload("Actor").
		Where("actor_id = ?", id).
		Find(&movies).Error

	if err != nil {
		return nil, err
	}

	return movies, nil
}
