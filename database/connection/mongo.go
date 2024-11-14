package connection

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type MongoDatabase struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func buildUri() string {
	uri := fmt.Sprintf("mongodb://%s:%s/",
		os.Getenv("DB_MONGO_HOST"),
		os.Getenv("DB_MONGO_PORT"),
	)

	return uri
}

func getDatabase() string {
	return os.Getenv("DB_MONGO_DATABASE")
}

func openConnectionMongo(uri string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// set client options
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	return client, nil
}

func pingConnectionMongo(database *mongo.Database) error {
	var result bson.M
	if err := database.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return fmt.Errorf("failed to ping the database: %w", err)
	}

	return nil
}

func closeConnectionMongo(client *mongo.Client) error {
	if err := client.Disconnect(context.TODO()); err != nil {
		return fmt.Errorf("failed to close the database: %w", err)
	}

	return nil
}

func ConnectMongo() (*MongoDatabase, error) {
	// open a connection to the database
	client, err := openConnectionMongo(buildUri())
	if err != nil {
		return nil, err
	}

	// connect to the database name
	database := client.Database(getDatabase())

	// ping the database to check the connection
	if err = pingConnectionMongo(database); err != nil {
		return nil, err
	}

	return &MongoDatabase{
		Client:   client,
		Database: database,
	}, nil
}

func (db *MongoDatabase) Close() error {
	// close a connection to the database
	if err := closeConnectionMongo(db.Client); err != nil {
		return fmt.Errorf("errors occurred while closing the databases: %v", err)
	}

	return nil
}
