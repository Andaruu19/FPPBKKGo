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
	MovieUsecase *usecases.MovieUsecase
	ActorUsecase *usecases.ActorUsecase
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

func (ac *AlbumController) CreateGeneratedAlbum(c *gin.Context) {
	// Struct untuk menangkap input JSON
	var albumRequest struct {
		Name string `json:"name" binding:"required"`
	}

	// Bind JSON dari request ke struct
	if err := c.ShouldBindJSON(&albumRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cari informasi aktor berdasarkan nama
	actor, err := ac.ActorUsecase.GetActorByName(albumRequest.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Actor not found"})
		return
	}

	// Buat slug unik berdasarkan ID aktor
	slug := fmt.Sprintf("actor-%d", actor.ID)

	// Buat deskripsi berdasarkan nama aktor
	deskripsi := fmt.Sprintf("Movie collection where %s is played", actor.Name)

	// Buat instance album
	album := domain.Album{
		Name:      actor.Name, // Nama aktor
		Slug:      slug,       // Slug unik
		Deskripsi: deskripsi,  // Deskripsi otomatis
		Generated: true,       // Tandai sebagai album yang dihasilkan
	}

	// Simpan album ke database
	createdAlbum, err := ac.AlbumUsecase.CreateGeneratedAlbum(album)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create album"})
		return
	}

	// Kembalikan respons sukses
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

func (ac *AlbumController) GetAllGeneratedAlbums(c *gin.Context) {
	albums, err := ac.AlbumUsecase.GetAllGeneratedAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func (ac *AlbumController) GetAlbumByID(c *gin.Context) {
	// Extract the album_id from the URL parameter
	idStr := c.Param("album_id")
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

	// Bind JSON payload to albumRequest
	if err := c.ShouldBindJSON(&albumRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Extract ID from URL parameter and convert it to uint
	idStr := c.Param("album_id")
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
	idStr := c.Param("album_id")
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

func (ac *AlbumController) AddMovieToAlbum(c *gin.Context) {
	// Parse album_id and movie_id from the URL parameters
	albumIDStr := c.Param("album_id")
	movieIDStr := c.Param("movie_id")

	albumID, err := strconv.Atoi(albumIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	// Call the usecase to add the movie to the album
	err = ac.AlbumUsecase.AddMovieToAlbum(uint(albumID), uint(movieID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie added to the album successfully"})
}

func (ac *AlbumController) GenerateMovieToAlbum(c *gin.Context) {
	// Parse album_id and movie_id from the URL parameters
	albumIDStr := c.Param("album_id")
	actorIDStr := c.Param("actor_id")

	albumID, err := strconv.Atoi(albumIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	actorID, err := strconv.Atoi(actorIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor ID"})
		return
	}

	movies, err := ac.MovieUsecase.GetMovieByActor(uint(actorID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	for _, movie := range movies {
		err = ac.AlbumUsecase.AddMovieToAlbum(uint(albumID), uint(movie.ID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movies added to the album successfully"})
}

func (ac *AlbumController) RemoveMovieFromAlbum(c *gin.Context) {
	// Parse album_id and movie_id from the URL parameters
	albumIDStr := c.Param("album_id")
	movieIDStr := c.Param("movie_id")

	albumID, err := strconv.Atoi(albumIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	// Call the usecase to remove the movie from the album
	err = ac.AlbumUsecase.RemoveMovieFromAlbum(uint(albumID), uint(movieID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie removed from the album successfully"})
}

// GetMoviesInAlbum is a method to fetch movies in a specific album
func (ac *AlbumController) GetMoviesInAlbum(c *gin.Context) {
	albumIDStr := c.Param("album_id")
	albumID, err := strconv.Atoi(albumIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	movies, err := ac.AlbumUsecase.GetMoviesInAlbum(uint(albumID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		return
	}

	c.JSON(http.StatusOK, movies)
}
