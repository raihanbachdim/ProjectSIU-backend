package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raihanbachdim/ProjectSIU/util"
)

func IsAuthenticate(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "tidak terautentikasi",
		})
	}
	return c.Next()
}
