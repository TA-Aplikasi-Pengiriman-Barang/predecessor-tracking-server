package infrastructure

import (
	"tracking-server/infrastructure/bus"
	"tracking-server/infrastructure/healthcheck"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	Healthcheck healthcheck.Controller
	Bus         bus.Controller
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewController); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck controller")
	}

	if err := container.Provide(bus.NewController); err != nil {
		return errors.Wrap(err, "failed to provide bus controller")
	}

	return nil
}

func Routes(app *fiber.App, controller Holder) {
	controller.Healthcheck.Routes(app)
	controller.Bus.Routes(app)
}
