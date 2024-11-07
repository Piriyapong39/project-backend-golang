package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/template/html/v2"
)

type Books struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Books

func main() {

	// append book
	books = append(books, Books{ID: 1, Title: "AjarnDaeng Guitar", Author: "AjarnDaeng"})
	books = append(books, Books{ID: 2, Title: "Dhama Chatri", Author: "Ajarn-Mai-Rom"})

	// config view
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// all rountes
	app.Get("/ping", greetUser)
	app.Get("/books", getBooks)
	app.Get("/books/:id", getOneBook)
	app.Post("/books", addBook)
	app.Patch("/books", updateBook)
	app.Delete("books", deleteBook)
	app.Post("/upload", uploadFile)
	app.Get("/test-html", testHtml)

	// running server on port
	app.Listen(":8080")
}
