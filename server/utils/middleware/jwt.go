package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"server/types/error_response"
	"server/types/secure"
	"server/utils/config"
)

func Jwt() fiber.Handler {
	conf := jwtware.Config{
		SigningKey:  []byte(config.C.JwtSecret),
		TokenLookup: "header:Homie",
		ContextKey:  "home",
		Claims:      new(secure.ClaimsStruct),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if len(c.Get("Homie")) == 0 {
				return &error_response.Error{
					Message: "JWT Validation error",
					Err:     err,
				}
			}

			//if strings.HasPrefix(c.Path(), "/api") {
			//	return c.Next()
			//}

			return &error_response.Error{
				Message: "JWT Validation error",
				Err:     err,
			}
		},
	}
	return jwtware.New(conf)
}
