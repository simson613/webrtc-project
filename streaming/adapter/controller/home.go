package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (ctl *Controller) HomeHandler(c *fiber.Ctx) error {
	return c.Render("home", nil, "layouts/main")
}
