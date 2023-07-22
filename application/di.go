package application

import (
	"students-tech-server/application/project"
	"students-tech-server/application/students"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	StudentsService students.Service
	ProjectService  project.Service
}

func Register(container *dig.Container) error {
	if err := container.Provide(students.NewService); err != nil {
		return errors.Wrap(err, "failed to provide student service")
	}
	if err := container.Provide(project.NewService); err != nil {
		return errors.Wrap(err, "failed to provide project service")
	}

	return nil
}
