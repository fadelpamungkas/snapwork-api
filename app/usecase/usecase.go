package usecase

import (
	"context"
	"golangapi/app/models"
)

type (
	UserUsecaseI interface {
		GetAllUC(ctx context.Context) (res models.UserResponse, err error)
		GetOneUserUC(ctx context.Context, id string) (res models.UserResponse, err error)
		InsertUC(ctx context.Context, req models.UserRequest) (res int, err error)
		UpdateRoleUC(ctx context.Context, req models.UserRoleRequest) (res int, err error)
		DeleteUC(ctx context.Context, id string) (res int, err error)

		LoginUC(ctx context.Context, req models.LoginRequest) (res models.UserResponse, err error)
	}

	PostUsecaseI interface {
		GetAllUC(ctx context.Context, query models.Query) (res models.PostResponse, err error)
		GetAllByUserUC(ctx context.Context, id string) (res models.PostResponse, err error)
		GetOneUC(ctx context.Context, id string) (res models.PostResponse, err error)
		InsertUC(ctx context.Context, req models.PostRequest) (res int, err error)
		UpdateUC(ctx context.Context, req models.PostRequest) (res int, err error)
		DeleteUC(ctx context.Context, id string) (res int, err error)
	}

	TransactionUsecaseI interface {
		InsertOrderUC(ctx context.Context, req models.OrderRequest) (res int, err error)
		GetAllOrderUC(ctx context.Context) (res models.OrderResponse, err error)
	}

	CompanyUsecaseI interface {
		InsertCompanyUC(ctx context.Context, req models.CompanyRequest) (res int, err error)
		GetAllCompaniesUC(ctx context.Context) (res models.CompanyResponse, err error)
		InsertJobUC(ctx context.Context, req models.CompanyJobRequest) (res int, err error)
		GetCompany(ctx context.Context, id string) (res models.CompanyResponse, err error)
		GetJobCompany(ctx context.Context, companyid string, jobId string) (res models.CompanyJobResponse, err error)
	}
)
