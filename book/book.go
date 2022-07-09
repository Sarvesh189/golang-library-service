package book

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	ISBN      int
	Title     string
	Publisher string
	Price     float64
}

// thsi is

func GetBooks() {
	//	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, e := mongo.Connect(ctx)
	if e != nil {
		log.Fatal(e)
	}

	e = client.Ping(context.TODO(), nil)
	if e != nil {
		log.Fatal(e)
	}

	collection := client.Database("libraryDB").Collection("books")
	//fmt.Println(collection)
	//var books []Book
	var bk Book

	exErr := collection.FindOne(ctx, bson.D{{"isbn", 1001}}).Decode(&bk)
	if exErr != nil {
		log.Fatal(exErr)
	}
	//cur.All(ctx, &books)
	fmt.Println(bk)
}
