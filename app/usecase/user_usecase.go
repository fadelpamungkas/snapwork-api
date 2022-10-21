package usecase

import (
	"context"
	"golangapi/app/models"
	"log"

	"golangapi/app/repository"
)

type UserUsecase struct {
	repo repository.UserRepositoryI
}

func NewUserUsecase(userRepo repository.UserRepositoryI) UserUsecaseI {
	return &UserUsecase{
		repo: userRepo,
	}
}

func (uc *UserUsecase) LoginUC(ctx context.Context, req models.LoginRequest) (res models.UserResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//login
	res, err = uc.repo.Login(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *UserUsecase) GetAllUC(ctx context.Context) (res models.UserResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//get all data
	list, err := uc.repo.GetAll(ctx)
	if err != nil {
		log.Println("failed to show data product with default log")
		return list, err
	}

	return list, err
}

func (uc *UserUsecase) GetOneUserUC(ctx context.Context, id string) (res models.UserResponse, err error) {
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

func (uc *UserUsecase) InsertUC(ctx context.Context, req models.UserRequest) (res int, err error) {
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

func (uc *UserUsecase) UpdateRoleUC(ctx context.Context, req models.UserRoleRequest) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	res, err = uc.repo.UpdateRole(ctx, req)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (uc *UserUsecase) DeleteUC(ctx context.Context, id string) (res int, err error) {
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
