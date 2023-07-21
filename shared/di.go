package shared

import (
	"students-tech-server/shared/config"
	"students-tech-server/shared/depedencies"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	Env  *config.Env
	Http *fiber.App
	DB   *pgxpool.Pool
}

func Register(container *dig.Container) error {
	if err := container.Provide(config.NewEnv); err != nil {
		return errors.Wrap(err, "failed to provide config")
	}

	if err := container.Provide(depedencies.NewHttp); err != nil {
		return errors.Wrap(err, "failed to provide http")
	}

	if err := container.Provide(depedencies.NewDatabase); err != nil {
		return errors.Wrap(err, "failed to provide database")
	}

	return nil
}
