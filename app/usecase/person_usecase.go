package usecase

import (
	"context"
	"golangapi/app/models"
	"log"

	"golangapi/app/repository"
)

type PersonUsecase struct {
	repo repository.PersonRepositoryI
}

func NewPersonUsecase(personRepo repository.PersonRepositoryI) PersonUsecaseI {
	return &PersonUsecase{
		repo: personRepo,
	}
}

func (uc *PersonUsecase) InsertPersonUC(ctx context.Context, req models.PersonRequest) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//insert data
	res, err = uc.repo.InsertPerson(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *PersonUsecase) GetPersonUC(ctx context.Context, id string) (res models.PersonResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//get data by id
	data, err := uc.repo.GetPerson(ctx, id)
	if err != nil {
		log.Println("failed to show data product with default log")
		return data, err
	}

	return data, err
}

func (uc *PersonUsecase) UpdatePersonUC(ctx context.Context, req models.PersonRequest) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	res, err = uc.repo.UpdatePerson(ctx, req)
	if err != nil {
		return res, err
	}
	return res, nil
}
