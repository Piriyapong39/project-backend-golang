package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/template/html/v2"

	"github.com/joho/godotenv"

	"github.com/aujito/managebook/middlewares"

	jwtware "github.com/gofiber/contrib/jwt"
)

type Books struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Books

func main() {

	// Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	JWT_TOKEN := os.Getenv("JWT_SECRET_KEY")

	// append book
	books = append(books, Books{ID: 1, Title: "AjarnDaeng Guitar", Author: "AjarnDaeng"})
	books = append(books, Books{ID: 2, Title: "Dhama Chatri", Author: "Ajarn-Mai-Rom"})

	// config view
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// use middlewares
	app.Use(middlewares.CheckMiddleware)

	// all rountes
	app.Post("/login", userLogin)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(JWT_TOKEN)},
	}))

	app.Get("/ping", greetUser)
	app.Get("/books", getBooks)
	app.Get("/books/:id", getOneBook)
	app.Post("/books", addBook)
	app.Patch("/books", updateBook)
	app.Delete("books", deleteBook)
	app.Post("/upload", uploadFile)
	app.Get("/test-html", testHtml)
	app.Get("/config", getEnv)

	// running server on ports
	app.Listen(":8080")
}
