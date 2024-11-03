package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Create book struct
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// create slice of book to psudo database
var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Guitar classic", Author: "Ajarn Daeng"})
	books = append(books, Book{ID: 2, Title: "Dhamma", Author: "Ajarn Beer"})

	// Get all book
	app.Get("/books", getBooks)
	// Get a single book
	app.Get("/books/:id", getBook)
	// Append book
	app.Post("/books", createBook)
	// Put book
	app.Put("/books/:id", updateBook)

	// Running server
	app.Listen(":8080")
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {

	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Cannot found book")
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpate := new(Book)
	if err := c.BodyParser(bookUpate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookUpate.Title
			books[i].Author = bookUpate.Author
			return c.JSON(books[i])
		}
	}
	return c.Status(fiber.StatusBadRequest).SendString("Book is not found")
}
