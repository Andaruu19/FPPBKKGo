package domain

import (
	"gorm.io/gorm"
)

type Genre struct {
    gorm.Model       // Includes ID, CreatedAt, UpdatedAt, DeletedAt fields
    Slug  string     `gorm:"type:varchar(255);unique;not null"`
    Name  string     `gorm:"type:varchar(255);not null"`
    Movies []Movie   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Relation with movies
}