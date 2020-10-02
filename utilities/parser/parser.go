package parser

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saifoelloh/minimalism_api/utilities/responses"
)

func BodyParser(ctx *fiber.Ctx, body interface{}) (err error) {
	if err := ctx.BodyParser(body); err != nil {
		return responses.ErrorResponse(ctx, fiber.StatusBadRequest, err)
	}
	return
}

func QueryParser(ctx *fiber.Ctx, query interface{}) (err error) {
	if err := ctx.QueryParser(query); err != nil {
		return responses.ErrorResponse(ctx, fiber.StatusBadRequest, err)
	}
	return
}
