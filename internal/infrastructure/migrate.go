package infrastructure

import (
	"FPPBKKGo/internal/infrastructure/database"
	"FPPBKKGo/internal/infrastructure/migrations"
	"log"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
    log.Println("Starting database migration...")

    // Call migration functions
    migrations.MigrateGenres(db)  // For genres
    migrations.MigrateActors(db)  // For actors
	database.SeedGenres(db)  // For genres
	database.SeedActors(db)  // For actors
    migrations.MigrateMovies(db)  // For movies
	database.SeedMovies(db)  // For movies
	database.SeedMovieGenresActors(db)  // For movie_genres and movie_actors

    log.Println("Database migration completed successfully.")
}
