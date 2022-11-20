package application

import (
	"tracking-server/application/bus"
	"tracking-server/application/healthcheck"
	"tracking-server/application/news"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthcheckService healthcheck.Service
	BusService         bus.Service
	NewsService        news.Service
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewHealthcheckService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck service")
	}

	if err := container.Provide(bus.NewBusService); err != nil {
		return errors.Wrap(err, "failed to provide bus service")
	}

	if err := container.Provide(news.NewNewsService); err != nil {
		return errors.Wrap(err, "failed to provide news service")
	}

	return nil
}
