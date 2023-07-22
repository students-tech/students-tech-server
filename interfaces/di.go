package interfaces

import (
	"students-tech-server/interfaces/health"
	"students-tech-server/interfaces/students"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthService   health.ViewService
	StudentsService students.ViewService
}

func Register(container *dig.Container) error {
	if err := container.Provide(health.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide health service")
	}

	if err := container.Provide(students.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide students service")
	}

	return nil
}
