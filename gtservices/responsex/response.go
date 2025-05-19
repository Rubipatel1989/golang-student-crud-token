package responsex

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GTSuccess(c *fiber.Ctx, success bool, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": success,
		"message": message,
		"data":    data,
	})
}

func GTError(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}

func BadRequest(c *fiber.Ctx, message string) error {
	return GTError(c, fiber.StatusBadRequest, message)
}

func Unauthorized(c *fiber.Ctx, message string) error {
	return GTError(c, fiber.StatusUnauthorized, strings.ToLower(message))
}
