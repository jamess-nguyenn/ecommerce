package mongo

import (
	"context"
	"ecommerce/database/connection"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository[model any] interface {
	Insert(document *model) (*mongo.InsertOneResult, error)
	InsertMany(documents []any) error
	// TransactionInsertMany use this function later
	TransactionInsertMany(documents []any) error
}

type BaseRepository[model any] struct {
	db   *connection.MongoDatabase
	name string
}

func NewBaseRepository[model any](db *connection.MongoDatabase, name string) *BaseRepository[model] {
	return &BaseRepository[model]{
		db:   db,
		name: name,
	}
}

func (repository *BaseRepository[model]) Name() string {
	return repository.name
}

func (repository *BaseRepository[model]) Insert(document *model) (*mongo.InsertOneResult, error) {
	result, err := repository.db.Database.Collection(repository.Name()).InsertOne(context.TODO(), document)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repository *BaseRepository[model]) InsertMany(documents []any) error {
	_, err := repository.db.Database.Collection(repository.Name()).InsertMany(context.TODO(), documents)
	if err != nil {
		return err
	}

	return nil
}

func (repository *BaseRepository[model]) TransactionInsertMany(documents []any) error {
	// start a session
	session, err := repository.db.Client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.TODO())

	// define the transactional operation
	callback := func(sessionContext mongo.SessionContext) (any, error) {
		result, err := repository.db.Database.Collection(repository.Name()).InsertMany(sessionContext, documents)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	// start the transaction with the callback
	_, err = session.WithTransaction(context.TODO(), callback)
	if err != nil {
		return err
	}

	return nil
}