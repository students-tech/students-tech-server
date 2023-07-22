package interfaces

import (
	"students-tech-server/interfaces/assignment"
	"students-tech-server/interfaces/health"
	"students-tech-server/interfaces/project"
	"students-tech-server/interfaces/students"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthService     health.ViewService
	ProjectService    project.ViewService
	StudentsService   students.ViewService
	AssignmentService assignment.ViewService
}

func Register(container *dig.Container) error {
	if err := container.Provide(health.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide health service")
	}

	if err := container.Provide(students.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide students service")
	}

	if err := container.Provide(project.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide students service")
	}

	if err := container.Provide(assignment.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide assignment service")
	}

	return nil
}
