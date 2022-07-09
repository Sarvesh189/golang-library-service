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

func getdbCollection() *mongo.Collection {
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

	return collection
}

// thsi is
func GetBookByISBN(isbn int) Book {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := getdbCollection()
	//fmt.Println(collection)
	//var books []Book
	var bk Book

	exErr := collection.FindOne(ctx, bson.D{{"isbn", 1001}}).Decode(&bk)
	if exErr != nil {
		log.Fatal(exErr)
	}
	//cur.All(ctx, &books)
	fmt.Println(bk)

	return bk
}

func GetAllBook() ([]Book, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	bookCollection := getdbCollection()
	//fmt.Println(collection)
	var books []Book
	var bk Book

	cursor, exErr := bookCollection.Find(ctx, bson.D{})
	if exErr != nil {
		defer cursor.Close(ctx)
		log.Fatal(exErr)
		return books, exErr
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(&bk)
		if err != nil {
			return books, err
		}
		books = append(books, bk)
	}
	return books, nil
}

func InsertBook(book Book) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	bookCollection := getdbCollection()
	rs, err := bookCollection.InsertOne(ctx, book)

	if err != nil {
		return "-1", err
	}
	return fmt.Sprintf("%v", rs.InsertedID), nil

}
