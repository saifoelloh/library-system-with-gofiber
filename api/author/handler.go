package author

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/saifoelloh/minimalism_api/utilities/database"
	"github.com/saifoelloh/minimalism_api/utilities/parser"
	"github.com/saifoelloh/minimalism_api/utilities/responses"
)

type Query struct {
	SortBy, OrderBy string
	Show, Page      int
}

func GetAuthors(ctx *fiber.Ctx) error {
	db := database.DB
	q := new(Query)
	var authors []Author

	parser.QueryParser(ctx, q)
	order := q.SortBy + " " + q.OrderBy
	offset := q.Page * q.Show

	if result := db.Preload("Books").Limit(q.Show).Offset(offset).Order(order).Find(&authors); result.Error != nil {
		return responses.ErrorResponse(ctx, fiber.StatusConflict, result.Error)
	}

	return responses.SuccessResponse(ctx, fiber.StatusOK, authors)
}

func CreateAuthor(ctx *fiber.Ctx) error {
	db := database.DB
	body := new(Author)

	parser.BodyParser(ctx, body)
	if result := db.Create(&body); result.Error != nil {
		return responses.ErrorResponse(ctx, fiber.StatusConflict, result.Error)
	}

	return responses.SuccessResponse(ctx, fiber.StatusCreated, body)
}

func GetAuthorById(ctx *fiber.Ctx) error {
	db := database.DB
	var author Author

	id, _ := strconv.Atoi(ctx.Params("id"))
	if result := db.Find(&author, id); result.Error != nil {
		return responses.ErrorResponse(ctx, fiber.StatusConflict, result.Error)
	}

	return responses.SuccessResponse(ctx, fiber.StatusOK, author)
}

func UpdateAuthorById(ctx *fiber.Ctx) error {
	db := database.DB
	author := new(Author)

	parser.BodyParser(ctx, author)
	if response := db.Model(&Author{}).Where("id", ctx.Params("id")).Updates(author); response.Error != nil {
		return responses.ErrorResponse(ctx, fiber.StatusConflict, response.Error)
	}

	return responses.SuccessResponse(ctx, fiber.StatusNoContent, nil)
}

func DeleteAuthorById(ctx *fiber.Ctx) error {
	db := database.DB
	id, _ := strconv.Atoi(ctx.Params("id"))

	if response := db.Delete(&Author{}, id); response.Error != nil {
		return responses.ErrorResponse(ctx, fiber.StatusConflict, response.Error)
	}

	return responses.SuccessResponse(ctx, fiber.StatusNoContent, nil)
}
