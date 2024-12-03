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
	dsn := "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	dbName := "fpgo"
	// Drop database if exist
	err = db.Exec("DROP DATABASE IF EXISTS " + dbName).Error
	if err != nil {
		log.Fatalf("Failed to drop database: %v", err)
	}
	log.Printf("Database %s dropped successfully.", dbName)

	// Create new database
	err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
	log.Printf("Database %s created successfully.", dbName)

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

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Set up routes
	routes.SetupMovieRoutes(router, movieController)
	routes.SetupAlbumRoutes(router, albumController)

	// Run the server
	log.Println("Server started on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
