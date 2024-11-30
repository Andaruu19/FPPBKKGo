package usecases

import (
	"FPPBKKGo/internal/domain"
	"errors"
	"strings"
	"unicode"

	"github.com/gosimple/slug"
)

// AlbumUsecase contains methods related to albums
type AlbumUsecase struct {
    AlbumRepository domain.AlbumRepository
	MovieRepository domain.MovieRepository
}

func (au *AlbumUsecase) CreateAlbum(album domain.Album) (domain.Album, error) {
    album.Slug = GenerateSlug(album.Name)

    // Save to repository
    createdAlbum, err := au.AlbumRepository.CreateAlbum(album)
    if err != nil {
        return domain.Album{}, err
    }

    return createdAlbum, nil
}

func (au *AlbumUsecase) GetAllAlbums() ([]domain.Album, error) {
    return au.AlbumRepository.GetAllAlbums()
}

func (au *AlbumUsecase) GetAlbumByID(id uint) (domain.Album, error) {
	return au.AlbumRepository.GetAlbumByID(id)
}

func (au *AlbumUsecase) UpdateAlbum(id uint, updatedAlbum domain.Album) (domain.Album, error) {
    // Fetch the existing album
    existingAlbum, err := au.AlbumRepository.GetAlbumByID(id)
    if err != nil {
        return domain.Album{}, err
    }

    // Update the slug only if the name is modified
    if existingAlbum.Name != updatedAlbum.Name {
        updatedAlbum.Slug = GenerateSlug(updatedAlbum.Name)
    }

    // Update in repository
    return au.AlbumRepository.UpdateAlbum(id, updatedAlbum)
}

func (au *AlbumUsecase) DeleteAlbum(id uint) error {
    // Ensure the album exists before attempting to delete
    _, err := au.AlbumRepository.GetAlbumByID(id)
    if err != nil {
        return errors.New("album not found")
    }

    // Call the repository to delete the album
    return au.AlbumRepository.DeleteAlbum(id)
}

func GenerateSlug(name string) string {
	// Remove non-alphanumeric characters except spaces
	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, name)

	return slug.Make(cleaned)
}

// AddMovieToAlbum adds a movie to the album
func (au *AlbumUsecase) AddMovieToAlbum(albumID uint, movieID uint) error {
    album, err := au.AlbumRepository.GetAlbumByID(albumID)
    if err != nil {
        return err
    }

    // Get the movie using MovieRepository
    movie, err := au.MovieRepository.GetByID(movieID)
    if err != nil {
        return err
    }

    // Check if the movie is already in the album
    for _, m := range album.Movies {
        if m.ID == movieID {
            return errors.New("movie already added to the album")
        }
    }

    // Add the movie to the album's Movies slice (this will reflect in the join table)
    album.Movies = append(album.Movies, movie)

    // Save the updated relationship to the database (this will update the join table)
    if err := au.AlbumRepository.AddMovieToAlbum(albumID, movieID); err != nil {
        return err
    }

    return nil
}

func (au *AlbumUsecase) RemoveMovieFromAlbum(albumID, movieID uint) error {
    // Use the repository method to remove the movie from the album_movies table
    if err := au.AlbumRepository.RemoveMovieFromAlbum(albumID, movieID); err != nil {
        return err
    }
    return nil
}

