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

func (as *DiplomserviceImpl) GetAllRestaraunt(ctx context.Context) ([]*model.RestarauntsModel, error) {
	rests, err := as.dipRepository.FindAllRestaraunts(ctx)
	if err != nil {
		return nil, err
	}
	return rests, nil
}
func (as *DiplomserviceImpl) NewRestaraunt(ctx context.Context, Name string, Category int) (string, error) {
	msg, err := as.dipRepository.NewRestaraunts(ctx, Name, Category)
	if err != nil {
		return msg, err
	}
	return msg, nil
}
func (as *DiplomserviceImpl) GetCattegories(ctx context.Context) ([]*model.Category, error) {
	categories, err := as.dipRepository.AllCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
