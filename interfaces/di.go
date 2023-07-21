package interfaces

import (
	"students-tech-server/interfaces/health"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthService health.ViewService
}

func Register(container *dig.Container) error {
	if err := container.Provide(health.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide healt service")
	}

	return nil
}
