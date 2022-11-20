package interfaces

import (
	"tracking-server/interfaces/bus"
	"tracking-server/interfaces/healthcheck"
	"tracking-server/interfaces/news"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthcheckViewService healthcheck.ViewService
	BusViewService         bus.ViewService
	NewsViewService        news.ViewService
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck view service")
	}

	if err := container.Provide(bus.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide bus view service")
	}

	if err := container.Provide(news.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide news view service")
	}

	return nil
}
