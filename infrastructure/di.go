package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/dig"
	"tracking-server/infrastructure/healthcheck"
)

type Holder struct {
	dig.In
	Healthcheck healthcheck.Controller
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewController); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck controller")
	}

	return nil
}

func Routes(app *fiber.App, controller Holder) {
	controller.Healthcheck.Routes(app)
}
