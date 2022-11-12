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
		UpdateRole(ctx context.Context, req models.UserRoleRequest) (res int, err error)
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
		InsertApplication(ctx context.Context, req models.ApplicationRequest) (res int, err error)
		GetAllApplication(ctx context.Context) (res models.ApplicationResponse, err error)
		GetAllApplicationByCompanyId(ctx context.Context, id string) (res models.ApplicationResponse, err error)
		GetAllApplicationByUserId(ctx context.Context, id string) (res models.ApplicationResponse, err error)
	}

	CompanyRepositoryI interface {
		InsertCompany(ctx context.Context, req models.CompanyRequest) (res int, err error)
		GetAllCompanies(ctx context.Context) (res models.CompanyResponse, err error)
		InsertJob(ctx context.Context, req models.CompanyJobRequest) (res int, err error)
		GetCompany(ctx context.Context, id string) (res models.CompanyResponse, err error)
		GetCompanyByUserId(ctx context.Context, id string) (res models.CompanyResponse, err error)
		GetJobCompany(ctx context.Context, companyId string, jobId string) (res models.CompanyJobResponse, err error)
		UpdateJobCompany(ctx context.Context, req models.CompanyJobRequest) (res int, err error)
		DeleteJobCompany(ctx context.Context, companyId string, jobId string) (res int, err error)
	}

	NewsRepositoryI interface {
		InsertNews(ctx context.Context, req models.NewsRequest) (res int, err error)
		GetAllNews(ctx context.Context) (res models.NewsResponse, err error)
		GetOneNews(ctx context.Context, id string) (res models.NewsResponse, err error)
		UpdateNews(ctx context.Context, req models.NewsRequest) (res int, err error)
		DeleteNews(ctx context.Context, id string) (res int, err error)
	}
)
