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
