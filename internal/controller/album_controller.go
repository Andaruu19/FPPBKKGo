package controllers

import (
	"FPPBKKGo/internal/domain"
	"FPPBKKGo/internal/usecases"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	AlbumUsecase *usecases.AlbumUsecase
}

func (ac *AlbumController) CreateAlbum(c *gin.Context) {
	var albumRequest struct {
		Name      string `json:"name" binding:"required"`
		Deskripsi string `json:"deskripsi" binding:"required"`
	}

	if err := c.ShouldBindJSON(&albumRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album := domain.Album{
		Name:      albumRequest.Name,
		Deskripsi: albumRequest.Deskripsi,
	}

	createdAlbum, err := ac.AlbumUsecase.CreateAlbum(album)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create album"})
		return
	}

	c.JSON(http.StatusCreated, createdAlbum)
}

func (ac *AlbumController) GetAllAlbums(c *gin.Context) {
	albums, err := ac.AlbumUsecase.GetAllAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func (ac *AlbumController) GetAlbumByID(c *gin.Context) {
    idStr := c.Param("id") // Extract the ID from the URL parameter
    id, err := strconv.Atoi(idStr) // Convert the string ID to an integer
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    // Fetch the album from the usecase
    album, err := ac.AlbumUsecase.GetAlbumByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
        return
    }

    c.JSON(http.StatusOK, album)
}

func (ac *AlbumController) UpdateAlbum(c *gin.Context) {
	var albumRequest struct {
		Name      string `json:"name" binding:"required"`
		Deskripsi string `json:"deskripsi" binding:"required"`
	}

	// Extract ID from URL parameter and convert it to uint
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert string to uint (32 bits)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	album := domain.Album{
		Name:      albumRequest.Name,
		Deskripsi: albumRequest.Deskripsi,
	}

	// Call the UpdateAlbum method of the usecase with the correct ID type (uint)
	updatedAlbum, err := ac.AlbumUsecase.UpdateAlbum(uint(id), album)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update album"})
		return
	}

	c.JSON(http.StatusOK, updatedAlbum)
}

func (ac *AlbumController) DeleteAlbum(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    fmt.Printf("Deleting album with ID: %d\n", id)

    err = ac.AlbumUsecase.DeleteAlbum(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}
