package usecase

import (
	"context"
	"golangapi/app/models"
	"log"

	"golangapi/app/repository"
)

type CompanyUsecase struct {
	repo repository.CompanyRepositoryI
}

func NewCompanyUsecase(companyRepo repository.CompanyRepositoryI) CompanyUsecaseI {
	return &CompanyUsecase{
		repo: companyRepo,
	}
}

func (uc *CompanyUsecase) InsertCompanyUC(ctx context.Context, req models.CompanyRequest) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//insert data
	res, err = uc.repo.InsertCompany(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *CompanyUsecase) UpdateCompanyStatusUC(ctx context.Context, req models.CompanyStatusRequest) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	res, err = uc.repo.UpdateCompanyStatus(ctx, req)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (uc *CompanyUsecase) GetAllCompaniesUC(ctx context.Context) (res models.CompanyResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//get all data
	list, err := uc.repo.GetAllCompanies(ctx)
	if err != nil {
		log.Println("failed to show data product with default log")
		return list, err
	}

	return list, err
}

func (uc *CompanyUsecase) InsertJobUC(ctx context.Context, req models.CompanyJobRequest) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//insert data
	res, err = uc.repo.InsertJob(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *CompanyUsecase) GetCompany(ctx context.Context, id string) (res models.CompanyResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//get data by id
	data, err := uc.repo.GetCompany(ctx, id)
	if err != nil {
		log.Println("failed to show data product with default log")
		return data, err
	}

	return data, err
}

func (uc *CompanyUsecase) GetCompanyByUserId(ctx context.Context, id string) (res models.CompanyResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//get data by id
	data, err := uc.repo.GetCompanyByUserId(ctx, id)
	if err != nil {
		log.Println("failed to show data product with default log")
		return data, err
	}

	return data, err
}

func (uc *CompanyUsecase) GetJobCompany(ctx context.Context, companyId string, jobId string) (res models.CompanyJobResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//get data by id
	data, err := uc.repo.GetJobCompany(ctx, companyId, jobId)
	if err != nil {
		log.Println("failed to show data product with default log")
		return data, err
	}

	return data, err
}

func (uc *CompanyUsecase) UpdateJobCompanyUC(ctx context.Context, req models.CompanyJobRequest) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	res, err = uc.repo.UpdateJobCompany(ctx, req)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (uc *CompanyUsecase) DeleteJobCompanyUC(ctx context.Context, companyId string, jobId string) (res int, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//delete data
	res, err = uc.repo.DeleteJobCompany(ctx, companyId, jobId)
	if err != nil {
		return res, err
	}

	return res, nil
}
