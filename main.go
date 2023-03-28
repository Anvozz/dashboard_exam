package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type User struct {
	ID    int
	Email string
}

func (u User) isAdministrator() bool {
	return true
}

func main() {
	engine := html.New("./views/src", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/index",
	})
	app.Static("/css", "./views/css")
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
			"User":  User{ID: 1, Email: "sankaapb@gmail.com"},
		})
	})

	app.Listen("localhost:3000")
}
