package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/saifoelloh/minimalism_api/utilities/database"
)

func GetUsers(ctx *fiber.Ctx) error {
	users := map[string]interface{}{}
	database.DB.Model(User{}).First(&users)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get all data",
		"status":  "ok",
		"payload": users,
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	customer := new(User)
	if err := ctx.BodyParser(customer); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	fmt.Println(customer)
	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message": "customer created",
		"payload": customer,
	})
}
