package service

import (
	"DiplomaWork/internal/app/model"
	"DiplomaWork/internal/app/repository"
	"context"
)

type DiplomaServiceImpl struct {
	dipRepository repository.DiplomaRepository
}

func NewDiplomaService(repo repository.DiplomaRepository) DiplomaService {
	return &DiplomaServiceImpl{dipRepository: repo}
}

func (as *DiplomaServiceImpl) GetAllRestaurants(ctx context.Context) (*model.Restaurant, error) {
	restaurants, err := as.dipRepository.FindAllRestaurants(ctx)
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}
