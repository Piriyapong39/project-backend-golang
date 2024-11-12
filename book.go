package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var memberUser = User{
	Email:    "example@example.com",
	Password: "password1234",
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

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := c.SaveFile(file, "./upload/"+file.Filename); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"msg": "Upload image successfully"})
}

func testHtml(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}

func getEnv(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"SECRET_KEY": os.Getenv("SECRET_KEY"),
	})
}

func userLogin(c *fiber.Ctx) error {
	JWT_TOKEN := os.Getenv("JWT_SECRET_KEY")
	fmt.Print(JWT_TOKEN)
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing Email or Password")
	}
	if user.Email != memberUser.Email || user.Password != memberUser.Password {
		return c.Status(fiber.StatusUnauthorized).SendString("Email or Password is invalid please check again")
	}
	claims := jwt.MapClaims{
		"email":    user.Email,
		"password": user.Password,
		"isAdmin":  "admin",
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(JWT_TOKEN))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg":   "Login successfully",
		"token": t,
	})
}
