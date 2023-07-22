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

// All godoc
// @Tags Healthcheck
// @Summary Check system status
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.Status
// @Failure 200 {array} dto.Status
// @Router /health [get]
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
