package domain

import (
	"gorm.io/gorm"
)

type Movie struct {
    gorm.Model
    Slug        string `gorm:"type:varchar(255);unique;not null"`
    Title       string `gorm:"type:varchar(255);not null"`
    Description string `gorm:"type:longtext;not null"`
    ImagePath   string `gorm:"type:varchar(255);not null"`
    Year        string `gorm:"type:varchar(255);not null"`
    GenreID     *uint  `gorm:"type:int;"` // Allow NULL for migration
    Genre       Genre  `gorm:"foreignKey:GenreID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
    ActorID     *uint  `gorm:"type:int;"` // Allow NULL for migration
    Actor       Actor  `gorm:"foreignKey:ActorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
    Albums      []*Album `gorm:"many2many:album_movies;constraint:OnDelete:CASCADE"`
}

type MovieRepository interface {
    GetByID(id uint) (*Movie, error)
    GetAll() ([]Movie, error)
    GetMoviesByAlbumID(albumID uint) ([]Movie, error) // New method to fetch movies by album ID
}