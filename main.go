package main

import (
	"fmt"

	"github.com/Sarvesh189/golang-library-service/route"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/template/html"
	// "github.com/markbates/goth"
	// "github.com/markbates/goth/gothic"
	// "github.com/markbates/goth/providers/github"
)

func main() {

	//goth.UseProviders()
	//key := "68abb8d6d7f3ecfc39f1"
	//secret := "Df9878cef0fce9117ce8ba4cd29e23580791ecf9"

	//github.New(key, secret, "http://localhost:3000/auth/github/callback")
	//fmt.Println("Server starting........../oaut")
	//engine := html.New("./login", ".html")

	//app := fiber.New(fiber.Config{Views: engine})
	app := fiber.New()
	route.CreateRoute(app)

	//app.Get("/", login)
	//book.GetBooks()

	//fmt.Println(bks)
	// app.Get("/auth/github/callback", func(c *fiber.Ctx) error {
	// 	user, err := gothic.CompleteUserAuth(ctx)
	// 	if err != nil {
	// 		return nil
	// 	}
	// 	fmt.Println(user.FirstName)

	// })

	app.Listen(":3000")

	fmt.Println("Server stoped", app)
}

func login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Hello, World!",
	})
}
