package health

import (
	"students-tech-server/interfaces"
	"students-tech-server/shared"
	"students-tech-server/shared/common"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Controller struct {
	Shared     shared.Holder
	Interfaces interfaces.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	app.Get("/health", c.healthcheck)
}

func (c *Controller) healthcheck(ctx *fiber.Ctx) error {
	log.Debug().Msg("checking server status")

	res, err := c.Interfaces.HealthService.SystemHealth()
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

func NewController(shared shared.Holder, interfaces interfaces.Holder) Controller {
	return Controller{
		Shared:     shared,
		Interfaces: interfaces,
	}
}
