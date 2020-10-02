package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saifoelloh/minimalism_api/api/author"
	"github.com/saifoelloh/minimalism_api/api/book"
)

func RouterSetup(app fiber.Router) {
	author.AuthorRouteSetup(app)
	book.BookRouterSetup(app)
}
