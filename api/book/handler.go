package book

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

func GetBooks(ctx *fiber.Ctx) error {
	db := database.DB
	q := new(Query)
	var books []Book

	parser.QueryParser(ctx, q)
	order := q.SortBy + " " + q.OrderBy
	offset := q.Page * q.Show
	if result := db.Joins("JOIN authors ON authors.id=books.author_id").Limit(q.Show).Offset(offset).Order(order).Find(&books); result.Error != nil {
		return responses.ErrorResponse(ctx, fiber.StatusConflict, result.Error)
	}

	return responses.SuccessResponse(ctx, fiber.StatusOK, books)
}

func CreateBook(ctx *fiber.Ctx) error {
	db := database.DB
	body := new(Book)

	parser.BodyParser(ctx, body)
	if result := db.Create(&body); result.Error != nil {
		return responses.ErrorResponse(ctx, fiber.StatusConflict, result.Error)
	}

	return responses.SuccessResponse(ctx, fiber.StatusCreated, body)
}

func GetBookById(ctx *fiber.Ctx) error {
	db := database.DB
	var book Book

	id, _ := strconv.Atoi(ctx.Params("id"))
	if result := db.Find(&book, id); result.Error != nil {
		return responses.ErrorResponse(ctx, fiber.StatusConflict, result.Error)
	}

	return responses.SuccessResponse(ctx, fiber.StatusOK, book)
}

func UpdateBookById(ctx *fiber.Ctx) error {
	db := database.DB
	book := new(Book)

	parser.BodyParser(ctx, book)
	if response := db.Model(&Book{}).Where("id", ctx.Params("id")).Updates(book); response.Error != nil {
		return responses.ErrorResponse(ctx, fiber.StatusConflict, response.Error)
	}

	return responses.SuccessResponse(ctx, fiber.StatusNoContent, nil)
}

func DeleteBookById(ctx *fiber.Ctx) error {
	db := database.DB
	id, _ := strconv.Atoi(ctx.Params("id"))

	if response := db.Delete(&Book{}, id); response.Error != nil {
		return responses.ErrorResponse(ctx, fiber.StatusConflict, response.Error)
	}

	return responses.SuccessResponse(ctx, fiber.StatusNoContent, nil)
}
