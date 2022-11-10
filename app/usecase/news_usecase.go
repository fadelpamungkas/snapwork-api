package usecase

import (
	"context"
	"golangapi/app/models"
	"log"

	"golangapi/app/repository"
)

type NewsUsecase struct {
	repo repository.NewsRepositoryI
}

func NewNewsUsecase(newsRepo repository.NewsRepositoryI) NewsUsecaseI {
	return &NewsUsecase{
		repo: newsRepo,
	}
}

func (uc *NewsUsecase) InsertNewsUC(ctx context.Context, req models.NewsRequest) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//insert data
	res, err = uc.repo.InsertNews(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}


func (uc *NewsUsecase) GetAllNewsUC(ctx context.Context) (res models.NewsResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//get all data
	list, err := uc.repo.GetAllNews(ctx)
	if err != nil {
		log.Println("failed to show data product with default log")
		return list, err
	}

	return list, err
}

func (uc *NewsUsecase) GetOneNewsUC(ctx context.Context, id string) (res models.NewsResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//get data by id
	data, err := uc.repo.GetOneNews(ctx, id)
	if err != nil {
		log.Println("failed to show data product with default log")
		return data, err
	}

	return data, err
}

func (uc *NewsUsecase) UpdateNewsUC(ctx context.Context, req models.NewsRequest) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	res, err = uc.repo.UpdateNews(ctx, req)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (uc *NewsUsecase) DeleteNewsUC(ctx context.Context, id string) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//delete data
	res, err = uc.repo.DeleteNews(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}
