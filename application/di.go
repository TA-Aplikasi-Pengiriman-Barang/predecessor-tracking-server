package application

import (
	"tracking-server/application/bus"
	"tracking-server/application/healthcheck"
	"tracking-server/application/news"
	"tracking-server/application/terminal"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthcheckService healthcheck.Service
	BusService         bus.Service
	NewsService        news.Service
	TerminalService    terminal.Service
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

	if err := container.Provide(terminal.NewTerminalService); err != nil {
		return errors.Wrap(err, "failed to provide terminal service")
	}

	return nil
}
