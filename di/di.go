package di

import (
	"students-tech-server/application"
	"students-tech-server/infrastructure"
	"students-tech-server/interfaces"
	"students-tech-server/shared"

	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

var Container = dig.New()

func init() {
	if err := shared.Register(Container); err != nil {
		log.Error().Err(err).Msg("error registering shared module")
	}

	if err := infrastructure.Register(Container); err != nil {
		log.Error().Err(err).Msg("error registering infrastructure module")
	}

	if err := interfaces.Register(Container); err != nil {
		log.Error().Err(err).Msg("error registering interfaces module")
	}

  if err := application.Register(Container); err != nil {
    log.Error().Err(err).Msg("error registering application module")
  }
}
