package main

import (
	"log"
	"students-tech-server/di"
	"students-tech-server/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func main() {
	container := di.Container

	err := container.Invoke(func(http *fiber.App, holder infrastructure.Holder) error {
		infrastructure.Routes(http, holder)

		err := http.Listen(":5000")
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
