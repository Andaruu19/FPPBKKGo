package usecases

import (
	"FPPBKKGo/internal/domain"
	repository "FPPBKKGo/internal/repositories/mysql"
	"fmt"
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

func (mu *MovieUsecase) GetMovieBySlug(slug string) (*domain.Movie, error) {
	if slug == "" {
		return nil, fmt.Errorf("slug cannot be empty")
	}

	movie, err := mu.MovieRepository.GetBySlug(slug)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (mu *MovieUsecase) GetMovieByName(name string) ([]domain.Movie, error) {
	if name == "" {
		return nil, fmt.Errorf("Movie name cannot be empty")
	}

	movie, err := mu.MovieRepository.GetByName(name)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (mu *MovieUsecase) GetMovieByGenre(slug string) ([]domain.Movie, error) {
	if slug == "" {
		return nil, fmt.Errorf("genre cannot be empty")
	}

	movie, err := mu.MovieRepository.GetByGenre(slug)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (mu *MovieUsecase) GetMovieByActor(id uint) ([]domain.Movie, error) {
	if id == 0 {
		return nil, fmt.Errorf("actor cannot be empty")
	}

	movie, err := mu.MovieRepository.GetByActor(id)
	if err != nil {
		return nil, err
	}

	return movie, nil
}
