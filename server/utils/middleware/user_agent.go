package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mileusna/useragent"
	"server/types/error_response"
	"server/utils/config"
)

func VerifyUserAgent(c *fiber.Ctx) error {
	ua := useragent.Parse(c.Get("User-Agent"))
	if ua.Name != config.C.UserAgent {
		return &error_response.Error{
			Message: "User-Agent is not valid",
		}
	}
	return c.Next()
}
