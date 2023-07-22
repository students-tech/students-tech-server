package infrastructure

import (
	"students-tech-server/infrastructure/assignment"
	"students-tech-server/infrastructure/health"
	"students-tech-server/infrastructure/project"
	"students-tech-server/infrastructure/students"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	Health     health.Controller
	Project    project.Controller
	Students   students.Controller
	Assignment assignment.Controller
}

func Register(container *dig.Container) error {
	if err := container.Provide(health.NewController); err != nil {
		return errors.Wrap(err, "failed to provide health controller")
	}

	if err := container.Provide(project.NewController); err != nil {
		return errors.Wrap(err, "failed to provide project controller")
	}
	if err := container.Provide(students.NewController); err != nil {
		return errors.Wrap(err, "failed to provide students controller")
	}
	if err := container.Provide(assignment.NewController); err != nil {
		return errors.Wrap(err, "failed to provide assignment controller")
	}

	return nil
}

func Routes(app *fiber.App, controller Holder) {
	controller.Health.Routes(app)
	controller.Students.Routes(app)
	controller.Project.Routes(app)
	controller.Assignment.Routes(app)

	log.Debug().Msg("load route complete")
}
