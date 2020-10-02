package user

import (
	"github.com/gofiber/fiber/v2"
)

func UserRouterSetup(app fiber.Router) {
	r := app.Group("/user")
	r.Get("/", GetUsers)
	r.Post("/", CreateUser)
}
