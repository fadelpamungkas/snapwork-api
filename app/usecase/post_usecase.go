package usecase

import (
	"context"
	"golangapi/app/models"
	"log"

	"golangapi/app/repository"
)

type PostUsecase struct {
	repo repository.PostRepositoryI
}

func NewPostUsecase(postRepo repository.PostRepositoryI) PostUsecaseI {
	return &PostUsecase{
		repo: postRepo,
	}
}

func (uc *PostUsecase) GetAllUC(ctx context.Context, query models.Query) (res models.PostResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//get all data
	list, err := uc.repo.GetAll(ctx, query)
	if err != nil {
		log.Println("failed to show data product with default log")
		return list, err
	}

	return list, err
}

func (uc *PostUsecase) GetAllByUserUC(ctx context.Context, id string) (res models.PostResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//get all data by user id
	list, err := uc.repo.GetAllByUser(ctx, id)
	if err != nil {
		log.Println("failed to show data product with default log")
		return list, err
	}

	return list, err
}

func (uc *PostUsecase) GetOneUC(ctx context.Context, id string) (res models.PostResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//get data by id
	data, err := uc.repo.GetOne(ctx, id)
	if err != nil {
		log.Println("failed to show data product with default log")
		return data, err
	}

	return data, err
}

func (uc *PostUsecase) InsertUC(ctx context.Context, req models.PostRequest) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//insert data
	res, err = uc.repo.Insert(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *PostUsecase) UpdateUC(ctx context.Context, req models.PostRequest) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	res, err = uc.repo.Update(ctx, req)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (uc *PostUsecase) DeleteUC(ctx context.Context, id string) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//delete data
	res, err = uc.repo.Delete(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}
