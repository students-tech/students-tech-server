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
	projectGroup := app.Group("/project")
	projectGroup.Post("/", c.createProject)
	projectGroup.Post("/list", c.listProject)
	projectGroup.Post("/detail", c.getProjectByID)
	projectGroup.Post("/completion", c.completionRequirements)
}

func (c *Controller) createProject(ctx *fiber.Ctx) error {
	log.Debug().Msg("start create project")

	var request dto.CreateProjectRequest
	if err := ctx.BodyParser(&request); err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}
	res, err := c.Interfaces.ProjectService.CreateProject(request)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

func (c *Controller) listProject(ctx *fiber.Ctx) error {
	log.Debug().Msg("start get list project")

	res, err := c.Interfaces.ProjectService.ListProject()
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

func (c *Controller) getProjectByID(ctx *fiber.Ctx) error {
	log.Debug().Msg("start get detail project")

	var request dto.GetProjectDetail
	if err := ctx.BodyParser(&request); err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}
	res, err := c.Interfaces.ProjectService.GetDetailProject(request.ProjectID)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}
func (c *Controller) completionRequirements(ctx *fiber.Ctx) error {
	log.Debug().Msg("start completion requirement")

	var request dto.CompleteRequirementRequest
	if err := ctx.BodyParser(&request); err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}
	resp, err := c.Interfaces.ProjectService.CompleteRequirement(request)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, resp)
}

func NewController(shared shared.Holder, interfaces interfaces.Holder) Controller {
	return Controller{
		Shared:     shared,
		Interfaces: interfaces,
	}
}
