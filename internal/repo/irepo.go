package repo

import (
	"context"
	"main/internal/entity"
)

type IContractRepo interface {
	CreateContract(ctx context.Context, contract *entity.Contract) error
	UpdateContract(ctx context.Context, studentCode string, contract *entity.Contract) error
	DeleteContract(ctx context.Context, studentCode string) error
	Search(ctx context.Context, keyword string) (entity.Contract, error)
	SearchAll() ([]entity.Contract, error)
}
