package author

import "github.com/gofiber/fiber/v2"

func AuthorRouteSetup(api fiber.Router) {
	r := api.Group("/author")
	r.Get("/", GetAuthors)
	r.Post("/", CreateAuthor)
	r.Get("/:id", GetAuthorById)
	r.Put("/:id", UpdateAuthorById)
	r.Delete("/:id", DeleteAuthorById)
}
