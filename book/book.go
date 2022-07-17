package book

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	libdbconnection "github.com/Sarvesh189/golang-library-service/dbconnection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBCollection struct {
	mongoCol *mongo.Collection
}

type Book struct {
	ISBN      int
	Title     string
	Publisher string
	Price     float64
}

var (
	dbConnOnce sync.Once
)

func getdbCollection() *mongo.Collection {
	var mcol *mongo.Collection

	dbConnOnce.Do(func() {
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		e := client.Connect(ctx)
		//	client, e := mongo.Connect(ctx)
		if e != nil {
			log.Fatal(e)
		}

		e = client.Ping(context.TODO(), nil)
		if e != nil {
			log.Fatal(e)
		}

		mcol = client.Database("libraryDB").Collection("books")
		//dbColl.mongoCol = collection

	})

	return mcol
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
	client, err := libdbconnection.GetMongoClient()
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	bookCollection := client.Database("libraryDB").Collection("books")
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
