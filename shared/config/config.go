package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Env struct {
	DBUrl          string `envconfig:"DB_URL" default:""`
	PORT           string `envconfig:"PORT" default:"8080"`
	HOST           string `envconfig:"HOST" default:"localhost:8080"`
	GCPProjectId   string `envconfig:"GCP_PROJECT_ID" default:""`
	GCPPubSubTopic string `envconfig:"GCP_PUB_SUB_TOPIC" default:""`
	ChatGPTHost    string `envconfig:"CHAT_GPT_HOST" default:""`
	ChatGPTAPIKey  string `envconfig:"CHAT_GPT_API_KEY" default:""`
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
