package main

import (
	"strconv"

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

	// greet user
	app.Get("/ping", greetUser)

	// get books data
	app.Get("/books", getBooks)

	// get one book by id
	app.Get("/books/:id", getOneBook)

	// Add book
	app.Post("/books", addBook)

	// running server on port
	app.Listen(":8080")
}

func greetUser(c *fiber.Ctx) error {
	return c.SendString("Pong")

}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getOneBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}

	return c.Status(fiber.StatusBadRequest).SendString("book is not found")
}

func addBook(c *fiber.Ctx) error {
	book := new(Books)
	err := c.BodyParser(book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if book.ID == 0 || book.Title == "" || book.Author == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing required fields")
	}

	for _, existBook := range books {
		if existBook.ID == book.ID {
			return c.Status(fiber.StatusBadRequest).SendString("Book id is already exists")
		}
	}

	books = append(books, *book)
	return nil
}
