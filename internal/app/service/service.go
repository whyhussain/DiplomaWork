package service

import (
	"DiplomaWork/internal/app/model"
	"context"
)

type DiplomaService interface {
	GetAllRestaraunt(ctx context.Context) (*model.RestarauntsModel, error)
}
