package usecase

import (
	"context"
	"golangapi/app/models"
	"log"

	"golangapi/app/repository"
)

type AssessmentUsecase struct {
	repo repository.AssessmentRepositoryI
}

func NewAssessmentUsecase(assessmentRepo repository.AssessmentRepositoryI) AssessmentUsecaseI {
	return &AssessmentUsecase{
		repo: assessmentRepo,
	}
}

func (uc *AssessmentUsecase) InsertAssessmentUC(ctx context.Context, req models.AssessmentRequest) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	res, err = uc.repo.InsertAssessment(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *AssessmentUsecase) GetAllAssessmentUC(ctx context.Context) (res models.AssessmentResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	list, err := uc.repo.GetAllAssessment(ctx)
	if err != nil {
		log.Println("failed to show data product with default log")
		return list, err
	}

	return list, err
}
