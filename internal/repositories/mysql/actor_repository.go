package repository

import (
	"FPPBKKGo/internal/domain"

	"gorm.io/gorm"
)

type ActorRepository struct {
	DB *gorm.DB
}

var _ domain.ActorRepository = (*ActorRepository)(nil)

func (ar *ActorRepository) GetActorByName(name string) (domain.Actor, error) {
	var actor domain.Actor
	if err := ar.DB.Where("actors.name = ?", name).First(&actor).Error; err != nil {
		return domain.Actor{}, err
	}
	return actor, nil
}
