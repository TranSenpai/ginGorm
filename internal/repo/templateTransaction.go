package repo

import (
	"context"

	"gorm.io/gorm"
)

type ITransaction interface {
	Execute(tx *gorm.DB) error
}

func RunTransaction(ctx context.Context, db *gorm.DB, t ITransaction) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return t.Execute(tx)
	})
}
