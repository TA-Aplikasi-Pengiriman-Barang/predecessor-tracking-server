package application

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
	"tracking-server/application/healthcheck"
)

type Holder struct {
	dig.In
	HealthcheckService healthcheck.Service
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewHealthcheckService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck service")
	}

	return nil
}
