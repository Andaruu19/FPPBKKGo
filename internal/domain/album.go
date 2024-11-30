package domain

import (
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
    Slug      string         `gorm:"type:varchar(255);uniqueIndex;not null"`
    Name      string         `gorm:"type:varchar(255);not null"`
    Deskripsi string         `gorm:"type:text;not null"`
    Movies    []*Movie       `gorm:"many2many:album_movies;"` // For many-to-many relationship
}