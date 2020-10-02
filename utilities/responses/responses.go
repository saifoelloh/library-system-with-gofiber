package responses

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseType struct {
	code            int
	status, message string
}

var (
	Success = map[int]ResponseType{
		200: {
			status:  "Success",
			message: "Success get your data.",
		},
		201: {
			status:  "created",
			message: "The request has been fulfilled, resulting in the creation of a new resource.",
		},
		204: {
			status:  "No Content",
			message: "The server successfully processed the request, and is not returning any content.",
		},
	}
	Error = map[int]ResponseType{
		400: {
			status:  "bad request",
			message: "bad request brother",
		},
		409: {
			status:  "Conflict",
			message: "Error while inserting data",
		},
	}
)

func SuccessResponse(ctx *fiber.Ctx, code int, data interface{}) error {
	resp := Success[code]
	return ctx.Status(code).JSON(&fiber.Map{
		"status":  resp.status,
		"message": resp.message,
		"payload": data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, code int, data interface{}) error {
	resp := Error[code]
	return ctx.Status(code).JSON(&fiber.Map{
		"status":  resp.status,
		"message": resp.message,
		"payload": data,
	})
}
