package usecases

import (
	"FPPBKKGo/internal/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// AlbumUsecase contains methods related to albums
type ActorUsecase struct {
	ActorRepository domain.ActorRepository
}

func (au *ActorUsecase) GetActorByName(name string) (*domain.Actor, error) {
	var actor domain.Actor
	actor, err := au.ActorRepository.GetActorByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("actor with name %s not found", name)
		}
		return nil, err
	}
	return &actor, nil
}
