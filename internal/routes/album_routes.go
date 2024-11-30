package routes

import (
	controllers "FPPBKKGo/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupAlbumRoutes(router *gin.Engine, albumController *controllers.AlbumController) {
	// Album routes - these deal with album information
	albumRoutes := router.Group("/albums")
	{
		albumRoutes.POST("/", albumController.CreateAlbum)        // Create a new album
		albumRoutes.GET("/", albumController.GetAllAlbums)        // Fetch all albums
		albumRoutes.GET("/:album_id", albumController.GetAlbumByID) // Get album by album_id
		albumRoutes.PUT("/:album_id", albumController.UpdateAlbum)  // Update album by album_id
		albumRoutes.DELETE("/:album_id", albumController.DeleteAlbum) // Delete album by album_id
	}

	// Movie routes for a specific album (no conflict as :album_id is distinct)
	movieRoutes := router.Group("/albums/:album_id/movies")
	{
		movieRoutes.POST("/:movie_id", albumController.AddMovieToAlbum)   // Add a movie to an album
		movieRoutes.DELETE("/:movie_id", albumController.RemoveMovieFromAlbum) // Remove a movie from an album
	}
}
