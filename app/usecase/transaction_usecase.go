package usecase

import (
	"context"
	"golangapi/app/models"
	"log"

	"golangapi/app/repository"
)

type TransactionUsecase struct {
	repo repository.TransactionRepositoryI
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepositoryI) TransactionUsecaseI {
	return &TransactionUsecase{
		repo: transactionRepo,
	}
}

func (uc *TransactionUsecase) InsertOrderUC(ctx context.Context, req models.OrderRequest) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//insert data
	res, err = uc.repo.InsertOrder(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}


func (uc *TransactionUsecase) GetAllOrderUC(ctx context.Context) (res models.OrderResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//get all data
	list, err := uc.repo.GetAllOrder(ctx)
	if err != nil {
		log.Println("failed to show data product with default log")
		return list, err
	}

	return list, err
}
