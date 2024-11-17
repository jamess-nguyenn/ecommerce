package mysql

import (
	"ecommerce/database/connection"
	"gorm.io/gorm"
)

type Creator[model any] interface {
	Create(record *model) (*model, error)
}

type Updater[model any] interface {
	Update(record *model) error
}

type BatchCreator[model any] interface {
	CreateMany(records []*model) error
}

type BatchUpdater[model any] interface {
	UpdateMany(records []*model) error
}

type Repository[model any] interface {
	Creator[model]
	Updater[model]
	BatchCreator[model]
	BatchUpdater[model]
}

type baseRepository[model any] struct {
	db *connection.MysqlDatabase
}

func newBaseRepository[model any](db *connection.MysqlDatabase) *baseRepository[model] {
	return &baseRepository[model]{db: db}
}

func (repository *baseRepository[model]) Create(record *model) (*model, error) {
	if err := repository.db.Master.Create(record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (repository *baseRepository[model]) CreateMany(records []*model) error {
	return repository.db.Master.Transaction(func(tran *gorm.DB) error {
		if err := tran.Create(records).Error; err != nil {
			return err
		}

		return nil
	})
}
