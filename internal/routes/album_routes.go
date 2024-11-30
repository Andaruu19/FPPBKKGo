package routes

import (
	controllers "FPPBKKGo/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupAlbumRoutes(router *gin.Engine, albumController *controllers.AlbumController) {
    albumRoutes := router.Group("/albums")
    {
        albumRoutes.POST("/", albumController.CreateAlbum)    // Create a new album
        albumRoutes.GET("/", albumController.GetAllAlbums)    // Fetch all albums
        albumRoutes.GET("/:id", albumController.GetAlbumByID) // Get album by ID
        albumRoutes.PUT("/:id", albumController.UpdateAlbum)  // Update album by ID
        albumRoutes.DELETE("/:id", albumController.DeleteAlbum) // Delete album by ID
    }
}