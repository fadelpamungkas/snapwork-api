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
		InsertApplicationUC(ctx context.Context, req models.ApplicationRequest) (res int, err error)
		UpdateApplicationStatusUC(ctx context.Context, req models.ApplicationStatusRequest) (res int, err error)
	}

	CompanyUsecaseI interface {
		InsertCompanyUC(ctx context.Context, req models.CompanyRequest) (res int, err error)
		UpdateCompanyStatusUC(ctx context.Context, req models.CompanyStatusRequest) (res int, err error)
		GetAllCompaniesUC(ctx context.Context) (res models.CompanyResponse, err error)
		InsertJobUC(ctx context.Context, req models.CompanyJobRequest) (res models.CompanyJobResponse, err error)
		GetCompanyUC(ctx context.Context, id string) (res models.CompanyResponse, err error)
		GetCompanyByUserIdUC(ctx context.Context, id string) (res models.CompanyResponse, err error)
		GetJobCompanyUC(ctx context.Context, companyid string, jobId string) (res models.CompanyJobResponse, err error)
		UpdateJobCompanyUC(ctx context.Context, req models.CompanyJobRequest) (res int, err error)
		DeleteJobCompanyUC(ctx context.Context, companyId string, jobId string) (res int, err error)
		UpdateJobPaymentUC(ctx context.Context, req models.CompanyJobPaymentRequest) (res int, err error)
	}

	NewsUsecaseI interface {
		InsertNewsUC(ctx context.Context, req models.NewsRequest) (res int, err error)
		GetAllNewsUC(ctx context.Context) (res models.NewsResponse, err error)
		GetOneNewsUC(ctx context.Context, id string) (res models.NewsResponse, err error)
		UpdateNewsUC(ctx context.Context, req models.NewsRequest) (res int, err error)
		DeleteNewsUC(ctx context.Context, id string) (res int, err error)
	}

	PersonUsecaseI interface {
		InsertPersonUC(ctx context.Context, req models.PersonRequest) (res int, err error)
		GetAllPersonUC(ctx context.Context) (res models.PersonResponse, err error)
		GetPersonUC(ctx context.Context, id string) (res models.PersonResponse, err error)
		UpdatePersonUC(ctx context.Context, req models.PersonRequest) (res int, err error)
		InsertNotificationUC(ctx context.Context, req models.Notification) (res int, err error)
		UpdateDocumentPersonUC(ctx context.Context, req models.PersonDocumentRequest) (res int, err error)
		UpdatePortfolioPersonUC(ctx context.Context, req models.Portfolio) (res int, err error)
		UpdateSelfDevelopmentPersonUC(ctx context.Context, req models.SelfDevelopment) (res int, err error)
		UpdateSelfDevelopmentPaymentPersonUC(ctx context.Context, req models.SelfDevelopmentPaymentRequest) (res int, err error)
	}

	AssessmentUsecaseI interface {
		InsertAssessmentUC(ctx context.Context, req models.AssessmentRequest) (res int, err error)
		GetAllAssessmentUC(ctx context.Context) (res models.AssessmentResponse, err error)
	}
)
