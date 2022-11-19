package application

import (
	"tracking-server/application/bus"
	"tracking-server/application/healthcheck"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthcheckService healthcheck.Service
	BusService         bus.Service
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewHealthcheckService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck service")
	}

	if err := container.Provide(bus.NewBusService); err != nil {
		return errors.Wrap(err, "failed to provide bus service")
	}

	return nil
}
