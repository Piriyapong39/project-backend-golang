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

	// all rountes
	app.Get("/ping", greetUser)
	app.Get("/books", getBooks)
	app.Get("/books/:id", getOneBook)
	app.Post("/books", addBook)
	app.Patch("/books", updateBook)

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

	for _, existBook := range books {
		if existBook.ID == book.ID {
			return c.Status(fiber.StatusBadRequest).SendString("This book id is already exists")
		}
	}

	books = append(books, *book)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg":  "Book added successfully",
		"data": book,
	})
}

func updateBook(c *fiber.Ctx) error {
	bookUpdate := new(Books)
	err := c.BodyParser(bookUpdate)
	print(err)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("eiei")
	}
	return nil
}
