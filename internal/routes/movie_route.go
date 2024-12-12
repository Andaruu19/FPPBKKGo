package routes

import (
	controllers "FPPBKKGo/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupMovieRoutes(router *gin.Engine, movieController *controllers.MovieController) {
	movieGroup := router.Group("/movies")
	{
		//movieGroup.GET("/:id", movieController.GetMovie)
		movieGroup.GET("/:slug", movieController.GetMovieBySlug)
		movieGroup.GET("/search/:title", movieController.GetMovieByTitle)
		movieGroup.GET("/", movieController.GetAllMovies)
		movieGroup.GET("/genre/:slug", movieController.GetMovieByGenre)
		movieGroup.GET("/actor/:id", movieController.GetMovieByActor)
	}
}
