package usecases

import (
	"FPPBKKGo/internal/domain"
	repository "FPPBKKGo/internal/repositories/mysql"
)

type MovieUsecase struct {
    MovieRepository *repository.MovieRepository
}

// GetMovie retrieves a movie by its ID
func (mu *MovieUsecase) GetMovie(id uint) (*domain.Movie, error) {
    movie, err := mu.MovieRepository.GetByID(id)
    if err != nil {
        return nil, err
    }
    return movie, nil
}

// GetAllMovies fetches all movies
func (mu *MovieUsecase) GetAllMovies() ([]domain.Movie, error) {
    movies, err := mu.MovieRepository.GetAll()
    if err != nil {
        return nil, err
    }
    return movies, nil
}
