package route

import (
	bookController "github.com/Sarvesh189/golang-library-service/book"
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
	return bookController.GetBookByISBN(c)

}

func getAllBook(c *fiber.Ctx) error {
	return bookController.GetBooks(c)
}

func insertBook(c *fiber.Ctx) error {
	return bookController.InsertBook(c)
}
