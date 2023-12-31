package main

import (
	"log"
	"students-tech-server/di"
	"students-tech-server/docs"
	"students-tech-server/infrastructure"
	"students-tech-server/shared/config"

	"github.com/gofiber/fiber/v2"
)

// @title Students-Tech Server
// @version 1.0
// @description API definition for Students Tech Server
// @host localhost:8080
// @BasePath /
func main() {
	container := di.Container

	err := container.Invoke(func(http *fiber.App, env *config.Env, holder infrastructure.Holder) error {
		infrastructure.Routes(http, holder)

		docs.SwaggerInfo.Host = env.HOST

		err := http.Listen(":" + env.PORT)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
