package assignment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"students-tech-server/interfaces"
	"students-tech-server/shared"
	"students-tech-server/shared/common"
	"students-tech-server/shared/dto"
)

type Controller struct {
	Shared     shared.Holder
	Interfaces interfaces.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	projectGroup := app.Group("/assignment")
	projectGroup.Post("/", c.createAssignment)
}

func (c *Controller) createAssignment(ctx *fiber.Ctx) error {
	log.Debug().Msg("start create assignment")

	var request dto.CreateAssignmentRequest
	if err := ctx.BodyParser(&request); err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}
	res, err := c.Interfaces.AssignmentService.CreateAssignment(request)
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
