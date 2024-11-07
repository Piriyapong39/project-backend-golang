package main

import (
	"github.com/gofiber/fiber/v2"
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
	app := fiber.New()

	// all rountes
	app.Get("/ping", greetUser)
	app.Get("/books", getBooks)
	app.Get("/books/:id", getOneBook)
	app.Post("/books", addBook)
	app.Patch("/books", updateBook)
	app.Delete("books", deleteBook)

	// running server on port
	app.Listen(":8080")
}
