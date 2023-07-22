package common

import (
	"context"
	"students-tech-server/shared/dto"

	"cloud.google.com/go/pubsub"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type (
	Response struct {
		Status string      `json:"status"`
		Error  string      `json:"error"`
		Data   interface{} `json:"data"`
	}
)

func DoCommonRequest(ctx *fiber.Ctx, body interface{}) error {
	err := ctx.BodyParser(body)
	if err != nil {
		log.Error().Err(err).Msg("error when parsing request body")
		return errors.New("failed to parse body")
	}

	validate := validator.New()
	err = validate.Struct(body)
	if err != nil {
		log.Error().Err(err).Msg("error when validating request body")
		return err
	}

	return nil
}

func DoCommonSuccessResponse(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(Response{
		Status: "SUCCESS",
		Data:   data,
	})
}

func DoCommonErrorResponse(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(Response{
		Status: "FAILED",
		Error:  err.Error(),
	})
}

func PublishEmailTopic(data dto.EmailMessage, projectId string, topic string) error {
	client, err := pubsub.NewClient(context.Background(), projectId)
	if err != nil {
		log.Error().Err(err).Msg("failed to initiate pub sub client")
		return err
	}

	topicPub := client.Topic(topic)

	dataByte, _ := json.Marshal(data)

	res := topicPub.Publish(context.Background(), &pubsub.Message{Data: dataByte})

	msgId, err := res.Get(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("failed to retrieve message id")
		return err
	}

	log.Debug().Str("id", msgId).Msg("pub sub message id")

	return nil
}
