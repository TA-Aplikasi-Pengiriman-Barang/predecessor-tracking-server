package interfaces

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
	"tracking-server/interfaces/healthcheck"
)

type Holder struct {
	dig.In
	HealthcheckViewService healthcheck.ViewService
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck view service")
	}

	return nil
}
