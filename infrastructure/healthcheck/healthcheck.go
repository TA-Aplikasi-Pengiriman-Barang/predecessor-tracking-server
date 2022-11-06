package healthcheck

import (
	"github.com/gofiber/fiber/v2"
	"tracking-server/interfaces"
	"tracking-server/shared"
)

type Controller struct {
	Interfaces interfaces.Holder
	Shared     shared.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	app.Get("/healthcheck", c.healthcheck)
}

func (c *Controller) healthcheck(ctx *fiber.Ctx) error {
	c.Shared.Logger.Println("checking server status")
	data, _ := c.Interfaces.HealthcheckViewService.SystemHealthcheck()
	return ctx.Status(fiber.StatusOK).JSON(data)
}

func NewController(interfaces interfaces.Holder, shared shared.Holder) Controller {
	return Controller{
		Interfaces: interfaces,
		Shared:     shared,
	}
}
