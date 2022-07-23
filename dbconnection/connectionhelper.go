package dbconnection

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var instanceError error
var mongoOnce sync.Once
var dbInstance *mongo.Database

const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DB               = "libraryDB"
	ISSUES           = "col_issues"
)

func GetMongoClient() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoOnce.Do(func() {
		clientSession, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
		if err != nil {
			instanceError = err
		}

		err = clientSession.Connect(ctx)
		if err != nil {
			instanceError = err
		}

		err = clientSession.Ping(ctx, nil)
		if err != nil {
			instanceError = err
		}

		clientInstance = clientSession
	})

	return clientInstance, instanceError
}
func GetMongoDBInstance() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoOnce.Do(func() {
		clientSession, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
		if err != nil {
			instanceError = err
		}

		err = clientSession.Connect(ctx)
		if err != nil {
			instanceError = err
		}

		err = clientSession.Ping(ctx, nil)
		if err != nil {
			instanceError = err
		}
		dbInstance = clientSession.Database(DB)

		//clientInstance = clientSession
	})

	return dbInstance, instanceError
}
