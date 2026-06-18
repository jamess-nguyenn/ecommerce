package mysql

import (
	"ecommerce/database/connection"
	"gorm.io/gorm"
)

// can delete these interfaces, but just keep it to implement this package to do with db
//type Creator[model any] interface {
//	Create(record *model) (*model, error)
//}
//
//type Updater[model any] interface {
//	Update(record *model) error
//}

type BatchCreator[model any] interface {
	CreateMany(records []*model) error
}

type BatchUpdater[model any] interface {
	UpdateMany(records []*model) error
}

//type Repository[model any] interface {
//	Creator[model]
//	Updater[model]
//	BatchCreator[model]
//	BatchUpdater[model]
//}

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

// check CreateInBatches here
func (repository *baseRepository[model]) CreateMany(records []*model) error {
	return repository.db.Master.Transaction(func(tran *gorm.DB) error {
		if err := tran.Create(records).Error; err != nil {
			return err
		}

		return nil
	})
}

func getColumns(columns []string) []string {
	if len(columns) == 0 {
		return []string{"*"}
	}

	return columns
}

//Create, Read, Update, Delete

func (repository *baseRepository[model]) buildQuery(conditions map[string]any, columns []string) *gorm.DB {
	return repository.db.Slave.Where(conditions).Select(getColumns(columns))
}

func (repository *baseRepository[model]) GetByPrimaryKey(id uint64, columns []string) (*model, error) {
	var record model
	conditions := map[string]any{"id": id}

	if err := repository.buildQuery(conditions, columns).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

func (repository *baseRepository[model]) GetByUniqueKey(conditions map[string]any, columns []string) (*model, error) {
	var record model

	if err := repository.buildQuery(conditions, columns).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

func (repository *baseRepository[model]) GetByColumns(conditions map[string]any, columns []string) ([]model, error) {
	var records []model

	if err := repository.buildQuery(conditions, columns).Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (repository *baseRepository[model]) UpdateById(id uint64, record *model) error {
	if err := repository.db.Master.Where("id = ?", id).Updates(&record).Error; err != nil {
		return err
	}

	return nil
}
