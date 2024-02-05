package db

import (
	"context"
	"fmt"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ERALiveDB *mongo.Database
)

func InitMongoDB(hostname string, username string, password string, dbName string) error {
	connectionStr := GetMongoDBConnectionString(hostname, username, password, dbName)
	logginghelper.LogInfo(connectionStr)
	client, err := GetMongoClient(connectionStr)
	if err != nil {
		return err
	}
	ERALiveDB = client.Database(dbName)
	return nil
}

func GetMongoClient(connectionString string) (*mongo.Client, error) {
	return mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
}

func GetMongoDBConnectionString(hostname string, username string, password string, dbName string) string {
	return fmt.Sprintf("mongodb://%s:%s@%s/%s", username, password, hostname, dbName)
}

func GetMongoCollection(collectionName string) *mongo.Collection {
	return ERALiveDB.Collection(collectionName)
}
