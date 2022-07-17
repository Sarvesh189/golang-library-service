package dbconnection

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

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
			clientInstanceError = err
		}

		err = clientSession.Connect(ctx)
		if err != nil {
			clientInstanceError = err
		}

		err = clientSession.Ping(ctx, nil)
		if err != nil {
			clientInstanceError = err
		}

		clientInstance = clientSession
	})

	return clientInstance, clientInstanceError
}
