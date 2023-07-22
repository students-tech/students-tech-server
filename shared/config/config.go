package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Env struct {
	DBUrl string `envconfig:"DB_URL" default:""`
	PORT  string `envconfig:"PORT" default:"8080"`
}

func NewEnv() (*Env, error) {
	var config Env

	_ = godotenv.Load()

	err := envconfig.Process("server", &config)
	if err != nil {
		log.Error().Err(err).Msg("error when processing environment")
		return nil, err
	}

	log.Info().Msg("processing environment success")

	return &config, nil
}
