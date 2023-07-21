package infrastructure

import (
	"students-tech-server/infrastructure/health"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	Health health.Controller
}

func Register(container *dig.Container) error {
	if err := container.Provide(health.NewController); err != nil {
		return errors.Wrap(err, "failed to provide health controller")
	}

	return nil
}

func Routes(app *fiber.App, controller Holder) {
	controller.Health.Routes(app)
	log.Debug().Msg("load route complete")
}
