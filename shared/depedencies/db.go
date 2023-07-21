package depedencies

import (
	"context"
	"students-tech-server/shared/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func NewDatabase(config *config.Env) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), config.DBUrl)
	if err != nil {
		log.Error().Err(err).Msg("error when connecting to database")
		return nil, err
	}

	log.Info().Msg("database connected")

	return pool, nil
}
