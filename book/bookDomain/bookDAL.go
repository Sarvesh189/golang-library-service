package bookDomain

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	libdbconnection "github.com/Sarvesh189/golang-library-service/dbconnection"
	"go.mongodb.org/mongo-driver/bson"
)

const collection_name = "books"

// thsi is
func getBookByISBN(isbn int) (*Book, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dbInstance, err := libdbconnection.GetMongoDBInstance()
	var bk Book

	if err != nil {
		book := new(Book)
		return book, nil
	}
	collection := dbInstance.Collection(collection_name)
	exErr := collection.FindOne(ctx, bson.D{{"isbn", isbn}}).Decode(&bk)
	if exErr != nil {
		log.Fatal(exErr)
	}
	log.Output(0, bk.Title+" "+strconv.Itoa(bk.ISBN))
	return &bk, nil
}

func getAllBook() ([]Book, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dbInstance, err := libdbconnection.GetMongoDBInstance()
	if err != nil {
		//fmt.Println(err)
		return nil, err
	}

	bookCollection := dbInstance.Collection(collection_name)
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

func insertBook(book Book) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dbInstance, err := libdbconnection.GetMongoDBInstance()
	if err != nil {
		return "", err
	}
	bookCollection := dbInstance.Collection(collection_name)
	rs, err := bookCollection.InsertOne(ctx, book)

	if err != nil {
		return "-1", err
	}
	return fmt.Sprintf("%v", rs.InsertedID), nil

}
