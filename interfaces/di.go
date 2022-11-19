package interfaces

import (
	"tracking-server/interfaces/bus"
	"tracking-server/interfaces/healthcheck"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthcheckViewService healthcheck.ViewService
	BusViewService         bus.ViewService
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck view service")
	}

	if err := container.Provide(bus.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide bus service")
	}

	return nil
}
