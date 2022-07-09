package main

import (
	"fmt"

	book "github.com/Sarvesh189/golang-library-service/book"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Server starting..........")
	app := fiber.New()

	book.GetBooks()

	//fmt.Println(bks)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, fiber")
	})

	app.Listen(":3000")

	fmt.Println("Server stoped", app)
}
