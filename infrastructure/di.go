package infrastructure

import (
	"tracking-server/infrastructure/bus"
	"tracking-server/infrastructure/healthcheck"
	"tracking-server/infrastructure/news"
	"tracking-server/infrastructure/terminal"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	Healthcheck healthcheck.Controller
	Bus         bus.Controller
	News        news.Controller
	Terminal    terminal.Controller
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewController); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck controller")
	}

	if err := container.Provide(bus.NewController); err != nil {
		return errors.Wrap(err, "failed to provide bus controller")
	}

	if err := container.Provide(news.NewController); err != nil {
		return errors.Wrap(err, "failed to provide news controller")
	}

	if err := container.Provide(terminal.NewController); err != nil {
		return errors.Wrap(err, "failed to provide terminal controller")
	}

	return nil
}

func Routes(app *fiber.App, controller Holder) {
	controller.Healthcheck.Routes(app)
	controller.Bus.Routes(app)
	controller.News.Routes(app)
	controller.Terminal.Routes(app)
}
