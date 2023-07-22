package project

import (
	"students-tech-server/interfaces"
	"students-tech-server/shared"
	"students-tech-server/shared/common"
	"students-tech-server/shared/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Controller struct {
	Shared     shared.Holder
	Interfaces interfaces.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	projectGroup := app.Group("/api/v1/project")
	projectGroup.Post("/", c.sendEmail)
	projectGroup.Post("/send-email", c.sendEmail)
}

func (c *Controller) createProject(ctx *fiber.Ctx) error {
	log.Debug().Msg("start create project")

	var request dto.SendEmailRequest
	if err := ctx.BodyParser(&request); err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}
	err := c.Interfaces.ProjectService.SendEmail(request)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, nil)
}

func (c *Controller) sendEmail(ctx *fiber.Ctx) error {
	log.Debug().Msg("start send email")

	var request dto.SendEmailRequest
	if err := ctx.BodyParser(&request); err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}
	err := c.Interfaces.ProjectService.SendEmail(request)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, nil)
}

func NewController(shared shared.Holder, interfaces interfaces.Holder) Controller {
	return Controller{
		Shared:     shared,
		Interfaces: interfaces,
	}
}
