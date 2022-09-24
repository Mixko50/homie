package fiber

import (
	"github.com/gofiber/fiber/v2"
	"server/types/error_response"
)

func notFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(error_response.Error{
		Success: false,
		Code:    "NOT_FOUND",
		Message: "Not found",
		Data:    nil,
	})
}
