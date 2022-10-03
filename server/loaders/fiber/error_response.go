package fiber

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"server/types/error_response"
	"strings"
)

func errorResponse(c *fiber.Ctx, err error) error {
	// * Case of *fiber.Error.
	if err, ok := err.(*fiber.Error); ok {
		return c.Status(err.Code).JSON(error_response.Error{
			Success: false,
			Code:    strings.ReplaceAll(strings.ToUpper(err.Error()), " ", "_"),
			Message: err.Error(),
			Data:    nil,
			Err:     err,
		})
	}

	// * Case of validator.ValidationErrors
	if e, ok := err.(validator.ValidationErrors); ok {
		var lists = make([]string, len(e))
		for i, err := range e {
			lists[i] = err.Error()
		}

		return c.Status(fiber.StatusBadRequest).JSON(error_response.Error{
			Success: false,
			Code:    "VALIDATION_ERROR",
			Message: "Validation error",
			Data:    lists,
			//Err:     err,
		})
	}

	if err, ok := err.(*error_response.Error); ok {
		// * Apply code fallback
		if err.Code == "" {
			err.Code = "GENERIC_ERROR"
		}

		return c.Status(fiber.StatusBadRequest).JSON(error_response.Error{
			Success: false,
			Code:    err.Code,
			Message: err.Message,
		})
	}
	log.Println(err.Error())

	return c.Status(fiber.StatusInternalServerError).JSON(error_response.Error{
		Success: false,
		Code:    "UNKNOWN_SERVER_SIDE_ERROR",
		Message: "Unknown server side error",
		Data:    nil,
		Err:     err,
	})
}
