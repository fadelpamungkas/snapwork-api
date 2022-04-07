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
		UpdateUC(ctx context.Context, req models.UserRequest) (res int, err error)
		DeleteUC(ctx context.Context, id string) (res int, err error)

		LoginUC(ctx context.Context, req models.LoginRequest) (res models.UserResponse, err error)
	}

	PostUsecaseI interface {
		GetAllUC(ctx context.Context) (res models.PostResponse, err error)
		GetAllByUserUC(ctx context.Context, id string) (res models.PostResponse, err error)
		GetOneUC(ctx context.Context, id string) (res models.PostResponse, err error)
		InsertUC(ctx context.Context, req models.PostRequest) (res int, err error)
		UpdateUC(ctx context.Context, req models.PostRequest) (res int, err error)
		DeleteUC(ctx context.Context, id string) (res int, err error)
	}
)
