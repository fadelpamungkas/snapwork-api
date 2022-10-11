package repository

import (
	"context"
	"golangapi/app/models"
)

type (
	UserRepositoryI interface {
		GetAll(ctx context.Context) (res models.UserResponse, err error)
		GetOne(ctx context.Context, id string) (res models.UserResponse, err error)
		Insert(ctx context.Context, req models.UserRequest) (res int, err error)
		Update(ctx context.Context, req models.UserRequest) (res int, err error)
		Delete(ctx context.Context, id string) (res int, err error)

		Login(ctx context.Context, req models.LoginRequest) (res models.UserResponse, err error)
	}

	PostRepositoryI interface {
		GetAll(ctx context.Context, query models.Query) (res models.PostResponse, err error)
		GetAllByUser(ctx context.Context, id string) (res models.PostResponse, err error)
		GetOne(ctx context.Context, id string) (res models.PostResponse, err error)
		Insert(ctx context.Context, req models.PostRequest) (res int, err error)
		Update(ctx context.Context, req models.PostRequest) (res int, err error)
		Delete(ctx context.Context, id string) (res int, err error)
	}

	TransactionRepositoryI interface {
		InsertOrder(ctx context.Context, req models.OrderRequest) (res int, err error)
		GetAllOrder(ctx context.Context) (res models.OrderResponse, err error)
	}

	CompanyRepositoryI interface {
		InsertCompany(ctx context.Context, req models.CompanyRequest) (res int, err error)
		GetAllCompanies(ctx context.Context) (res models.CompanyResponse, err error)
		InsertJob(ctx context.Context, req models.CompanyJobRequest) (res int, err error)
		GetCompany(ctx context.Context, id string) (res models.CompanyResponse, err error)
		GetJobCompany(ctx context.Context, companyId string, jobId string) (res models.CompanyJobResponse, err error)
	}
)
