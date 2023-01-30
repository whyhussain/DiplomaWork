package service

import (
	"DiplomaWork/internal/app/model"
	"DiplomaWork/internal/app/repository"
	"context"
)

type DiplomserviceImpl struct {
	dipRepository repository.DiplomaRepository
}

func NewDiplomaService(repo repository.DiplomaRepository) DiplomaService {
	return &DiplomserviceImpl{dipRepository: repo}
}

func (as *DiplomserviceImpl) GetAllRestaraunt(ctx context.Context) (*model.RestarauntsModel, error) {
	as.dipRepository.FindAllRestaraunts(ctx)
	return nil, nil
}
