package main

import (
	controllers "FPPBKKGo/internal/controller"
	"FPPBKKGo/internal/infrastructure"
	repositories "FPPBKKGo/internal/repositories/mysql"
	"FPPBKKGo/internal/routes"
	"FPPBKKGo/internal/usecases"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    // Connect to MySQL server without specifying a database
    dsn := "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Create the database
    dbName := "fpgo"
    err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error
    if err != nil {
        log.Fatalf("Failed to create database: %v", err)
    }

    // Connect to the newly created database
    dsn = "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Run database migrations
    infrastructure.MigrateDB(db)

    // Initialize repositories
    movieRepo := &repositories.MovieRepository{DB: db}
    albumRepo := &repositories.AlbumRepository{DB: db}

    // Initialize use cases
    movieUsecase := &usecases.MovieUsecase{MovieRepository: movieRepo}
    albumUsecase := &usecases.AlbumUsecase{
        AlbumRepository: albumRepo,
        MovieRepository: movieRepo,
    }

    // Initialize controllers
    movieController := &controllers.MovieController{MovieUsecase: movieUsecase}
    albumController := &controllers.AlbumController{AlbumUsecase: albumUsecase}

    // Set up Gin router
    router := gin.Default()

    // Set up routes
    routes.SetupMovieRoutes(router, movieController)
    routes.SetupAlbumRoutes(router, albumController)

    // Run the server
    log.Println("Server started on port 8080...")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start the server: %v", err)
    }
}
