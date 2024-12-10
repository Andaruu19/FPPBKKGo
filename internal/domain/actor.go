package domain

import (
	"gorm.io/gorm"
)

type Actor struct {
	gorm.Model
	Name   string  `gorm:"type:varchar(255);not null"`
	Movies []Movie `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type ActorRepository interface {
	GetActorByName(name string) (Actor, error)
}
