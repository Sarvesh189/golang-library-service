package book

import (
	"log"
	"strconv"

	bookDomain "github.com/Sarvesh189/golang-library-service/book/bookDomain"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	bks, err := bookDomain.GetAllBook()
	if err != nil {
		err, _ := err.(*bookDomain.BookError)

		return c.Status(err.Code).SendString(err.Err.Error())
	} else {
		return c.Status(200).JSON(bks)
	}

}
func GetBookByISBN(c *fiber.Ctx) error {
	isbn := c.Params("isbn")
	isbnN, _ := strconv.Atoi(isbn)
	bks, err := bookDomain.GetBookByISBN(isbnN)
	if err != nil {
		err, _ := err.(*bookDomain.BookError)
		return c.Status(err.Code).SendString(err.Err.Error())
	} else {
		return c.Status(200).JSON(bks)
	}

}

func InsertBook(c *fiber.Ctx) error {
	bk := bookDomain.Book{}

	err := c.BodyParser(bk)
	if err != nil {
		return c.Status(500).SendString("Parsing error")
	}
	log.Output(0, bk.Title)
	var bkEntity bookDomain.Book
	bkEntity.ISBN = bk.ISBN
	bkEntity.Title = bk.Title
	bkEntity.Publisher = bk.Publisher
	bkEntity.Price = bk.Price

	id, err := bookDomain.InsertBook(bkEntity)

	if err != nil {
		err, _ := err.(*bookDomain.BookError)
		return c.Status(err.Code).SendString(err.Err.Error())
	}
	return c.Status(200).SendString(id)

}
