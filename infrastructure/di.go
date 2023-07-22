package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
	"students-tech-server/infrastructure/health"
	"students-tech-server/infrastructure/project"
)

type Holder struct {
	dig.In
	Health  health.Controller
	Project project.Controller
}

func Register(container *dig.Container) error {
	if err := container.Provide(health.NewController); err != nil {
		return errors.Wrap(err, "failed to provide health controller")
	}

	if err := container.Provide(project.NewController); err != nil {
		return errors.Wrap(err, "failed to provide project controller")
	}

	return nil
}

func Routes(app *fiber.App, controller Holder) {
	controller.Health.Routes(app)
	controller.Project.Routes(app)
	log.Debug().Msg("load route complete")
}
