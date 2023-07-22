package depedencies

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"students-tech-server/shared/config"
)

func NewDatabase(config *config.Env) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), config.DBUrl)
	if err != nil {
		log.Error().Err(err).Msg("error when connecting to database")
		return nil, err
	}

	log.Info().Msg("database connected")

	//MigrateDatabase(config.DBUrl)

	return pool, nil
}

func MigrateDatabase(url string) {
	m, err := migrate.New("file://migrations", url)
	if err != nil {
		log.Warn().Err(err).Msg("error migrating database")
	}
	m.Up()
	log.Info().Msg("success migrating database")
}
