package shared

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"tracking-server/shared/config"
	"tracking-server/shared/depedencies"
)

type Holder struct {
	dig.In
	Logger *logrus.Logger
	Env    *config.EnvConfig
	Http   *fiber.App
}

func Register(container *dig.Container) error {
	if err := container.Provide(depedencies.NewLogger); err != nil {
		return errors.Wrap(err, "failed to provide logger")
	}

	if err := container.Provide(config.NewEnvConfig); err != nil {
		return errors.Wrap(err, "failed to provide env")
	}

	if err := container.Provide(depedencies.NewHttp); err != nil {
		return errors.Wrap(err, "failed to provide http")
	}

	return nil
}
