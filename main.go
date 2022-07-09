package main

import (
	"fmt"

	route "github.com/Sarvesh189/golang-library-service/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Server starting..........")
	app := fiber.New()
	route.CreateRoute(app)
	//	book.GetBooks()

	//fmt.Println(bks)

	//app.Get("/", func(c *fiber.Ctx) error {
	//return c.SendString("Hello, fiber")
	//})

	app.Listen(":3000")

	fmt.Println("Server stoped", app)
}
