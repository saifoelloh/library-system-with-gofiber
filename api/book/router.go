package book

import "github.com/gofiber/fiber/v2"

func BookRouterSetup(api fiber.Router) {
	r := api.Group("/book")
	r.Get("/", GetBooks)
	r.Post("/", CreateBook)
	r.Get("/:id", GetBookById)
	r.Put("/:id", UpdateBookById)
	r.Delete("/:id", DeleteBookById)
}
