package repo

import (
	"context"
	"main/internal/entity"
	"main/internal/models"
)

type IRepo interface {
	CreateContract(ctx context.Context, contract *entity.Contract) error
	UpdateContract(ctx context.Context, keyword any, contract *entity.Contract) error
	DeleteContract(ctx context.Context, keyword any) error
	Search(ctx context.Context, filter *models.Filter) ([]entity.Contract, error)
}
