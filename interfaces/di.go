package interfaces

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
	"students-tech-server/interfaces/health"
	"students-tech-server/interfaces/project"
)

type Holder struct {
	dig.In
	HealthService  health.ViewService
	ProjectService project.ViewService
}

func Register(container *dig.Container) error {
	if err := container.Provide(health.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide health service")
	}
	if err := container.Provide(project.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide project service")
	}

	return nil
}
