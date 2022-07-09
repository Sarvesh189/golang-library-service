package route

import (
	"fmt"
	"strconv"

	book "github.com/Sarvesh189/golang-library-service/book"
	"github.com/gofiber/fiber/v2"
)

type BookModel struct {
	ISBN      string
	Title     string
	Publisher string
	Price     string
}

func CreateRoute(app *fiber.App) {
	app.Get("/books/:isbn", getBookByISBN)
	app.Get("/books", getAllBook)
	app.Post("/books", insertBook)

}

func getBookByISBN(c *fiber.Ctx) error {
	isbn := c.Params("isbn")
	isbnN, _ := strconv.Atoi(isbn)
	bk := book.GetBookByISBN(isbnN)
	return c.Status(200).JSON(bk)
}

func getAllBook(c *fiber.Ctx) error {
	bks, err := book.GetAllBook()
	if err != nil {
		return c.Status(500).SendString("There is some server error. Please try later.")
	}
	return c.Status(200).JSON(bks)
}

func insertBook(c *fiber.Ctx) error {
	bk := &BookModel{}

	err := c.BodyParser(bk)
	fmt.Println(bk)
	var bkEntity book.Book
	bkEntity.ISBN, _ = strconv.Atoi(bk.ISBN)
	bkEntity.Title = bk.Title
	bkEntity.Publisher = bk.Publisher
	bkEntity.Price, _ = strconv.ParseFloat(bk.Price, 64)

	if err != nil {
		return c.Status(500).SendString("Parsing error")
	}
	id, err := book.InsertBook(bkEntity)
	if err != nil {
		return c.Status(500).SendString("Insert db error")
	}
	return c.Status(200).SendString(id)

}
