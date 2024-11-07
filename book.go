package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

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
	if book.ID == 0 || book.Author == "" || book.Title == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing some fields please check again")
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
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if bookUpdate.ID == 0 || bookUpdate.Author == "" || bookUpdate.Title == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing some fields")
	}
	for i, book := range books {
		if book.ID == bookUpdate.ID {
			books[i].Author = bookUpdate.Author
			books[i].Title = bookUpdate.Title
			return c.JSON(books[i])
		}
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "book is not found", "data": ""})
}

func deleteBook(c *fiber.Ctx) error {
	book := new(Books)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if book.ID == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Book ID is missing")
	}
	for i, bookDel := range books {
		if bookDel.ID == book.ID {
			books = append(books[:i], books[(i+1):]...)
			return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"msg": "remove book successfully", "data": bookDel})
		}
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Book is not found"})
}
