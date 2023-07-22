package main

import (
	"log"
	"students-tech-server/di"
	"students-tech-server/infrastructure"
	"students-tech-server/shared/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	container := di.Container

	err := container.Invoke(func(http *fiber.App, env *config.Env, holder infrastructure.Holder) error {
		infrastructure.Routes(http, holder)

    err := http.Listen(":8080")
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
